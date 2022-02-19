package kresty

import (
	"github.com/guojiangli/picasso/pkg/client/kresty/middleware"
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
	"github.com/guojiangli/picasso/pkg/trace/kjaeger"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go"
)

type Client struct {
	*resty.Client
	Logger baselogger.Logger
}

func New(opts ...*Option) *Client {
	khttpopts := defaultOption().MergeOption(opts...)
	cli := Client{
		Client: resty.NewWithClient(khttpopts.HTTPClient),
		Logger: khttpopts.Logger,
	}
	cli.SetTimeout(khttpopts.Timeout)
	cli.JSONMarshal = jsoniter.Marshal
	cli.JSONUnmarshal = jsoniter.Unmarshal
	return &cli
}

func (c *Client) WithSimpleLog() *Client {
	c.Client = c.OnAfterResponse(middleware.SimpleLog(c.Logger))
	return c
}

func (c *Client) WithTotalLog() *Client {
	c.Client = c.OnAfterResponse(middleware.TotalLog(c.Logger))
	return c
}

func (c *Client) WithKlogging() *Client {
	c.Client = c.OnAfterResponse(klog.RestyLogging())
	return c
}

func (c *Client) WithKTraceLog(l *klog.Logger) *Client {
	c.Client = c.OnBeforeRequest(klog.WithRestyTraceLog(l))
	return c
}

func (c *Client) WithMetric() *Client {
	c.Client = c.OnAfterResponse(middleware.Metric())
	return c
}

func (c *Client) WithTrace(tracer opentracing.Tracer) *Client {
	c.Client = c.OnBeforeRequest(kjaeger.RestyTracingBefore(tracer))
	c.Client = c.OnAfterResponse(kjaeger.RestyTracingAfter())
	return c
}
