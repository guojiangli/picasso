package middleware

import (
	"strings"

	"github.com/guojiangli/picasso/pkg/kmetric"
	"github.com/go-resty/resty/v2"
)

func Metric() func(c *resty.Client, res *resty.Response) error {
	return func(c *resty.Client, res *resty.Response) error {
		req := res.Request
		name := strings.Split(req.URL, "?")[0]
		kmetric.ClientHandleHistogram.Observe(res.Time().Seconds(), kmetric.TypeHTTP, name, req.Method)
		kmetric.ClientHandleCounter.Inc(kmetric.TypeHTTP, name, req.Method, res.Status())
		return nil
	}
}
