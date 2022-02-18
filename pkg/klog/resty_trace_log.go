package klog

import (
	"io/ioutil"

	"github.com/go-resty/resty/v2"
)

type logData struct {
	RequestHeaders  map[string][]string `json:"request_headers,omitempty"`
	ResponseHeaders map[string][]string `json:"response_headers,omitempty"`
	URI             string              `json:"uri"`
	Host            string              `json:"host,omitempty"`
	Status          string              `json:"status"`
	Method          string              `json:"method"`
	RequestBody     string              `json:"request_body,omitempty"`
	Proto           string              `json:"proto,omitempty"`
	ResponseBody    string              `json:"response_body,omitempty"`
	Latency         int64               `json:"latency(μs)"`
}

func RestyLogging() func(c *resty.Client, res *resty.Response) error {
	return func(_ *resty.Client, res *resty.Response) error {
		req := res.Request
		reqr := req.RawRequest
		resr := res.RawResponse
		var body []byte
		var err error
		if reqr.Body != nil {
			body, err = ioutil.ReadAll(reqr.Body)
			if err != nil {
				return err
			}
		}
		data := logData{
			Method:          req.Method,
			URI:             reqr.URL.RequestURI(),
			Proto:           reqr.Proto,
			Host:            reqr.URL.Host,
			Status:          res.Status(),
			Latency:         res.Time().Microseconds(),
			RequestHeaders:  reqr.Header,
			RequestBody:     string(body),
			ResponseHeaders: resr.Header,
			ResponseBody:    string(res.Body()),
		}
		FromTraceCtx(req.Context()).Info().Interface("details", data).Send()
		return nil
	}
}

func WithRestyTraceLog(l *Logger) func(*resty.Client, *resty.Request) error {
	return func(_ *resty.Client, r *resty.Request) error {
		c := r.Context()
		ctx := WithTraceCtx(l, c)
		r.SetContext(ctx)
		return nil
	}
}
