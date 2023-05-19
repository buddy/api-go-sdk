package buddy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func NewLoggingHttpTransport(t http.RoundTripper) *LoggingHttpTransport {
	return &LoggingHttpTransport{transport: t}
}

type LoggingHttpTransport struct {
	transport http.RoundTripper
}

func (t *LoggingHttpTransport) shouldLog() bool {
	tfLog := os.Getenv("TF_LOG")
	return tfLog == "DEBUG" || tfLog == "INFO" || tfLog == "TRACE"
}

func (t *LoggingHttpTransport) Log(msg string, a ...any) {
	if !t.shouldLog() {
		return
	}
	fmt.Printf(msg+"\n", a...)
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
