package kjaeger

import (
	"github.com/guojiangli/picasso/pkg/config"
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
	"github.com/opentracing/opentracing-go"
)

// Config HTTP config
type Option struct {
	Logger                   baselogger.JaegerLogger
	ServiceName              string
	SampleType               string
	LocalAgentHostPort       string
	TraceBaggageHeaderPrefix string
	TraceContextHeaderName   string
	BufferFlushInterval      string
	tags                     []opentracing.Tag
	SampleParam              float64
	EnableRPCMetrics         bool
	ReporterLogSpans         bool
}

func NewOption() *Option {
	return new(Option)
}

// RawConfig ...
func ConfigOption(key string) (option *Option, err error) {
	option = NewOption()
	err = config.UnmarshalKey(key, option)
	if err != nil {
		return nil, err
	}
	return option, nil
}

func (c *Option) SetServiceName(s string) *Option {
	c.ServiceName = s
	return c
}

func (c *Option) SetSampleType(l string) *Option {
	c.SampleType = l
	return c
}

func (c *Option) SetSampleParam(l float64) *Option {
	c.SampleParam = l
	return c
}

func (c *Option) SetReporterLogSpans(l bool) *Option {
	c.ReporterLogSpans = l
	return c
}

func (c *Option) SetBufferFlushInterval(l string) *Option {
	c.BufferFlushInterval = l
	return c
}

func (c *Option) SetLocalAgentHostPort(l string) *Option {
	c.LocalAgentHostPort = l
	return c
}

func (c *Option) SetTraceBaggageHeaderPrefix(l string) *Option {
	c.TraceBaggageHeaderPrefix = l
	return c
}

func (c *Option) SetTraceContextHeaderName(l string) *Option {
	c.TraceContextHeaderName = l
	return c
}

func (c *Option) SetPanicOnError(l string) *Option {
	c.TraceBaggageHeaderPrefix = l
	return c
}

func (c *Option) SetLogger(l baselogger.JaegerLogger) *Option {
	c.Logger = l
	return c
}

func (c *Option) SetTag(l ...opentracing.Tag) *Option {
	if c.tags == nil {
		c.tags = make([]opentracing.Tag, 0)
	}
	c.tags = append(c.tags, l...)
	return c
}

func SetServiceName(s string) *Option {
	c := NewOption()
	c.ServiceName = s
	return c
}

func SetSampleType(l string) *Option {
	c := NewOption()
	c.SampleType = l
	return c
}

func SetSampleParam(l float64) *Option {
	c := NewOption()
	c.SampleParam = l
	return c
}

func SetReporterLogSpans(l bool) *Option {
	c := NewOption()
	c.ReporterLogSpans = l
	return c
}

func SetBufferFlushInterval(l string) *Option {
	c := NewOption()
	c.BufferFlushInterval = l
	return c
}

func SetLocalAgentHostPort(l string) *Option {
	c := NewOption()
	c.LocalAgentHostPort = l
	return c
}

func SetTraceBaggageHeaderPrefix(l string) *Option {
	c := NewOption()
	c.TraceBaggageHeaderPrefix = l
	return c
}

func SetTraceContextHeaderName(l string) *Option {
	c := NewOption()
	c.TraceContextHeaderName = l
	return c
}

func SetPanicOnError(l string) *Option {
	c := NewOption()
	c.TraceBaggageHeaderPrefix = l
	return c
}

func SetLogger(l baselogger.JaegerLogger) *Option {
	c := NewOption()
	c.Logger = l
	return c
}

func SetTag(l ...opentracing.Tag) *Option {
	c := NewOption()
	c.tags = append(c.tags, l...)
	return c
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
		if opt.ServiceName != "" {
			c.ServiceName = opt.ServiceName
		}
		if opt.SampleType != "" {
			c.SampleType = opt.SampleType
		}
		if opt.SampleParam != 0 {
			c.SampleParam = opt.SampleParam
		}
		if opt.ReporterLogSpans {
			c.ReporterLogSpans = opt.ReporterLogSpans
		}
		if opt.BufferFlushInterval != "" {
			c.BufferFlushInterval = opt.BufferFlushInterval
		}
		if opt.TraceBaggageHeaderPrefix != "" {
			c.TraceBaggageHeaderPrefix = opt.TraceBaggageHeaderPrefix
		}
		if opt.TraceContextHeaderName != "" {
			c.TraceContextHeaderName = opt.TraceContextHeaderName
		}
		if opt.EnableRPCMetrics {
			c.EnableRPCMetrics = opt.EnableRPCMetrics
		}
		if opt.tags != nil {
			c.tags = opt.tags
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("kepler", "jaeger").Logger()

	return &Option{
		SampleType:               "const",
		SampleParam:              1,
		ReporterLogSpans:         true,
		BufferFlushInterval:      "1s",
		TraceBaggageHeaderPrefix: "ctx-",
		TraceContextHeaderName:   "x-trace-id",
		EnableRPCMetrics:         false,
		tags:                     nil,
		Logger: &klog.JaegerLogger{
			Logger: klog.NewLoggerWith(&l),
		},
	}
}
