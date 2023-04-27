package buddy

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"golang.org/x/time/rate"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultBaseUrl     = "https://api.buddy.works"
	headerRateRemainig = "X-Rate-Limit-Remaining"
	headerRateReset    = "X-Rate-Limit-Reset"
)

type UrlPath struct {
	Path   string
	Params []interface{}
}

func (u *UrlPath) Compute() string {
	return fmt.Sprintf(u.Path, u.Params...)
}

type Client struct {
	client *retryablehttp.Client

	baseUrl *url.URL

	mu      sync.Mutex
	limiter *rate.Limiter

	token string

	ProfileService       *ProfileService
	ProfileEmailService  *ProfileEmailService
	GroupService         *GroupService
	MemberService        *MemberService
	PermissionService    *PermissionService
	WorkspaceService     *WorkspaceService
	PublicKeyService     *PublicKeyService
	ProjectService       *ProjectService
	ProjectMemberService *ProjectMemberService
	ProjectGroupService  *ProjectGroupService
	WebhookService       *WebhookService
	VariableService      *VariableService
	SsoService           *SsoService
	IntegrationService   *IntegrationService
	PipelineService      *PipelineService
	TokenService         *TokenService
	SourceService        *SourceService
}

type PageQuery struct {
	Page    int `url:"page"`
	PerPage int `url:"per_page"`
}

func (c *Client) setBaseUrl(urlStr string) error {
	baseUrl, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	c.baseUrl = baseUrl
	return nil
}

func (c *Client) getLimiter() *rate.Limiter {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.limiter == nil {
		c.limiter = rate.NewLimiter(rate.Every(time.Second), 1000)
	}
	return c.limiter
}

func (c *Client) setLimiter(r rate.Limit, b int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.limiter.SetLimit(r)
	c.limiter.SetBurst(b)
}

func (c *Client) syncLimiter(r *http.Response) {
	headerReset := r.Header.Get(headerRateReset)
	headerRemaining := r.Header.Get(headerRateRemainig)
	if headerReset != "" && headerRemaining != "" {
		until, err1 := strconv.ParseInt(headerReset, 10, 64)
		remaining64, err2 := strconv.ParseInt(headerRemaining, 10, 64)
		if err1 == nil && err2 == nil {
			remaining := int(remaining64)
			seconds := int(until - time.Now().Unix())
			var b int
			var r rate.Limit
			if remaining <= 100 {
				b = 1
				r = rate.Limit(float64(remaining) / float64(seconds))
			} else {
				b = remaining - 100
				r = rate.Limit(1)
			}
			c.setLimiter(r, b)
		}
	}
}

func rateLimitBackoff(min, max time.Duration, resp *http.Response) time.Duration {
	// rnd is used to generate pseudo-random numbers.
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// First create some jitter bounded by the min and max durations.
	jitter := time.Duration(rnd.Float64() * float64(max-min))

	if resp != nil {
		if v := resp.Header.Get(headerRateReset); v != "" {
			if reset, _ := strconv.ParseInt(v, 10, 64); reset > 0 {
				// Only update min if the given time to wait is longer.
				if wait := time.Until(time.Unix(reset, 0)); wait > min {
					min = wait
				}
			}
		}
	}
	return min + jitter
}

func (c *Client) retryHTTPCheck(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 429 || resp.StatusCode >= 500 {
		return true, nil
	}
	return false, nil
}

func (c *Client) retryHTTPBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// Use the rate limit backoff function when we are rate limited.
	if resp != nil && resp.StatusCode == 429 {
		return rateLimitBackoff(min, max, resp)
	}
	// Set custom duration's when we experience a service interruption.
	min = 700 * time.Millisecond
	max = 1000 * time.Millisecond

	return retryablehttp.LinearJitterBackoff(min, max, attemptNum, resp)
}

func NewDefaultClient(token string) (*Client, error) {
	return NewClient(token, "", false)
}

func NewClient(token string, baseUrl string, insecure bool) (*Client, error) {
	tlsConfig := &tls.Config{}
	// turn off ssl verification
	if insecure {
		tlsConfig.InsecureSkipVerify = true
	}
	// configure transport
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.TLSClientConfig = tlsConfig
	t.MaxIdleConnsPerHost = 100
	// http client
	h := &http.Client{
		Transport: logging.NewTransport("Buddy", t),
		Timeout:   30 * time.Second,
	}
	// api client
	c := &Client{}
	if baseUrl != "" {
		err := c.setBaseUrl(baseUrl)
		if err != nil {
			return nil, err
		}
	} else {
		_ = c.setBaseUrl(defaultBaseUrl)
	}
	c.token = token
	c.client = &retryablehttp.Client{
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		Backoff:      c.retryHTTPBackoff,
		CheckRetry:   c.retryHTTPCheck,
		HTTPClient:   h,
		RetryMax:     5,
	}

	c.ProfileService = &ProfileService{client: c}
	c.ProfileEmailService = &ProfileEmailService{client: c}
	c.GroupService = &GroupService{client: c}
	c.MemberService = &MemberService{client: c}
	c.PermissionService = &PermissionService{client: c}
	c.WorkspaceService = &WorkspaceService{client: c}
	c.PublicKeyService = &PublicKeyService{client: c}
	c.ProjectService = &ProjectService{client: c}
	c.ProjectMemberService = &ProjectMemberService{client: c}
	c.ProjectGroupService = &ProjectGroupService{client: c}
	c.WebhookService = &WebhookService{client: c}
	c.VariableService = &VariableService{client: c}
	c.IntegrationService = &IntegrationService{client: c}
	c.PipelineService = &PipelineService{client: c}
	c.SourceService = &SourceService{client: c}
	c.SsoService = &SsoService{client: c}
	c.TokenService = &TokenService{client: c}
	return c, nil
}

func (c *Client) Create(url *UrlPath, postBody interface{}, query interface{}, respBody interface{}) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, url.Compute(), postBody, query)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if !c.isNil(respBody) {
		result = &respBody
	}
	return c.Do(req, result)
}

func (c *Client) Get(url *UrlPath, respBody interface{}, query interface{}) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodGet, url.Compute(), nil, query)
	if err != nil {
		return nil, err
	}
	return c.Do(req, &respBody)
}

func (c *Client) Patch(url *UrlPath, patchBody interface{}, query interface{}, respBody interface{}) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodPatch, url.Compute(), patchBody, query)
	if err != nil {
		return nil, err
	}
	return c.Do(req, &respBody)
}

func (c *Client) Put(url *UrlPath, putBody interface{}, query interface{}, respBody interface{}) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodPut, url.Compute(), putBody, query)
	if err != nil {
		return nil, err
	}
	return c.Do(req, &respBody)
}

func (c *Client) Delete(url *UrlPath, body interface{}, query interface{}) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodDelete, url.Compute(), body, query)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != http.StatusNoContent {
		return resp, errors.New("something went wrong while deleteing resource")
	}
	return resp, nil
}

func (c *Client) NewUrlPath(path string, a ...interface{}) *UrlPath {
	u := UrlPath{
		Path:   path,
		Params: a,
	}
	return &u
}

func (c *Client) isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func (c *Client) NewRequest(method, path string, body interface{}, params interface{}) (*retryablehttp.Request, error) {
	u := *c.baseUrl
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	// Set the encoded path data
	u.RawPath = c.baseUrl.Path + path
	u.Path = c.baseUrl.Path + unescaped
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Authorization", "Bearer "+c.token)
	var b interface{}
	if !c.isNil(body) {
		reqHeaders.Set("Content-Type", "application/json")
		b, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	if !c.isNil(params) {
		q, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}
	req, err := retryablehttp.NewRequest(method, u.String(), b)
	if err != nil {
		return nil, err
	}
	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}
	return req, nil
}

func (c *Client) Do(req *retryablehttp.Request, v interface{}) (*http.Response, error) {
	err := c.getLimiter().Wait(req.Context())
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	c.syncLimiter(res)
	err = CheckResponse(req, res)
	if err != nil {
		return res, err
	}
	if !c.isNil(v) {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
		}
	}
	return res, err
}

type ErrorResponse struct {
	Body     []byte
	Response *http.Response
	Request  *retryablehttp.Request
	Message  string
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Request.URL.Scheme, e.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d\n%s", e.Request.Method, u, e.Response.StatusCode, e.Message)
}

func CheckResponse(req *retryablehttp.Request, res *http.Response) error {
	switch res.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}
	errorResponse := &ErrorResponse{
		Response: res,
		Request:  req,
	}
	data, err := ioutil.ReadAll(res.Body)
	if err == nil && data != nil {
		errorResponse.Body = data
		var raw interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			errorResponse.Message = "failed to parse unknown error format"
		} else {
			errorResponse.Message = parseError(raw)
		}
	}
	return errorResponse
}

// Format:
// {
//     "errors": [
//			{
//				"message": "..."
//			}
//    	]
// }
func parseError(raw interface{}) string {
	switch raw := raw.(type) {
	case string:
		return raw
	case []interface{}:
		var errs []string
		for _, v := range raw {
			errs = append(errs, parseError(v))
		}
		return strings.Join(errs, "\n")
	case map[string]interface{}:
		var errs []string
		for _, v := range raw {
			errs = append(errs, parseError(v))
		}
		return strings.Join(errs, "\n")

	default:
		return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
	}
}
