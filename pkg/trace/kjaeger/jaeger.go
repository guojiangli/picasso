package kjaeger

import (
	"io"

	"github.com/guojiangli/picasso/pkg/utils/ktime"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func New(opts ...*Option) (opentracing.Tracer, io.Closer, error) {
	opt := defaultOption().MergeOption(opts...)
	configuration := config.Configuration{
		ServiceName: opt.ServiceName,
		RPCMetrics:  opt.EnableRPCMetrics,
		Tags:        opt.tags,
		Sampler: &config.SamplerConfig{
			Type:  opt.SampleType,
			Param: opt.SampleParam,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: ktime.Duration(opt.BufferFlushInterval),
			LogSpans:            opt.ReporterLogSpans,
			LocalAgentHostPort:  opt.LocalAgentHostPort,
		},
		Headers: &jaeger.HeadersConfig{
			TraceContextHeaderName:   opt.TraceContextHeaderName,
			TraceBaggageHeaderPrefix: opt.TraceBaggageHeaderPrefix,
		},
	}
	configuration.Sampler.Options = append(configuration.Sampler.Options, jaeger.SamplerOptions.Logger(opt.Logger))
	return configuration.NewTracer()
}
