package buddy

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// LogFunc receives a single formatted log message.
//
// A caller embedding the SDK in a host that owns the process output - a Terraform
// provider for instance - must route this to the host's own logger. Writing to
// stdout or stderr from inside such a host corrupts its plugin protocol.
type LogFunc func(msg string)

func NewLoggingHttpTransport(t http.RoundTripper) *LoggingHttpTransport {
	return &LoggingHttpTransport{transport: t}
}

type LoggingHttpTransport struct {
	transport http.RoundTripper
	log       LogFunc
}

// SetLogger installs the sink that requests and responses are reported to.
// Nothing is logged until one is set.
func (t *LoggingHttpTransport) SetLogger(log LogFunc) {
	t.log = log
}

func (t *LoggingHttpTransport) shouldLog() bool {
	return t.log != nil
}

func (t *LoggingHttpTransport) Log(msg string) {
	if !t.shouldLog() {
		return
	}
	t.log(msg)
}

func (t *LoggingHttpTransport) LogReq(req *http.Request) {
	if !t.shouldLog() {
		return
	}
	l := "API Request "
	l += req.Method
	l += " " + req.URL.RequestURI()
	if req.Body != nil {
		reqBody, err := io.ReadAll(req.Body)
		if err == nil && reqBody != nil {
			l += "\n" + t.formatJson(reqBody)
			req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}
	}
	t.Log(l)
}

func (t *LoggingHttpTransport) formatJson(rawJson []byte) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, rawJson, "", "\t")
	if err != nil {
		return string(rawJson)
	}
	return prettyJSON.String()
}

func (t *LoggingHttpTransport) LogRes(res *http.Response) {
	if !t.shouldLog() {
		return
	}
	l := "API Response "
	l += res.Request.Method
	l += " " + res.Request.URL.RequestURI()
	l += " " + res.Status
	if res.Body != nil {
		resBody, err := io.ReadAll(res.Body)
		if err == nil && resBody != nil {
			l += "\n" + t.formatJson(resBody)
			res.Body = io.NopCloser(bytes.NewBuffer(resBody))
		}
	}
	t.Log(l)
}

func (t *LoggingHttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.LogReq(req)
	res, err := t.transport.RoundTrip(req)
	if err != nil {
		return res, err
	}
	t.LogRes(res)
	return res, nil
}
