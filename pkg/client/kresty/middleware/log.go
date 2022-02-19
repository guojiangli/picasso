package middleware

import (
	"io/ioutil"

	"github.com/guojiangli/picasso/pkg/klog/baselogger"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type logData struct {
	Method          string              `json:"method"`
	URI             string              `json:"uri"`
	Proto           string              `json:"proto,omitempty"`
	Host            string              `json:"host,omitempty"`
	Status          string              `json:"status"`
	Latency         int64               `json:"latency(μs)"`
	RequestHeaders  map[string][]string `json:"request_headers,omitempty"`
	RequestBody     string              `json:"request_body,omitempty"`
	ResponseHeaders map[string][]string `json:"response_headers,omitempty"`
	ResponseBody    string              `json:"response_body,omitempty"`
}

func TotalLog(logger baselogger.Logger) func(c *resty.Client, res *resty.Response) error {
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
		Data := logData{
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
		s, err := jsoniter.Marshal(Data)
		if err != nil {
			return err
		}
		logger.Log(string(s))
		return nil
	}
}

func SimpleLog(logger baselogger.Logger) func(c *resty.Client, res *resty.Response) error {
	return func(_ *resty.Client, res *resty.Response) error {
		req := res.Request
		reqr := req.RawRequest
		Data := logData{
			Method:  req.Method,
			URI:     reqr.URL.RequestURI(),
			Status:  res.Status(),
			Latency: res.Time().Microseconds(),
		}
		s, err := jsoniter.Marshal(Data)
		if err != nil {
			return err
		}
		logger.Log(string(s))
		return nil
	}
}
