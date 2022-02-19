package kjaeger

import (
	"context"
	"strings"

	"github.com/guojiangli/picasso/pkg/trace/kopentracing"
	"github.com/go-resty/resty/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func RestyTracingAfter() func(*resty.Client, *resty.Response) error {
	return func(_ *resty.Client, r *resty.Response) error {
		ctx := r.Request.Context()
		Span := opentracing.SpanFromContext(ctx)
		Span.Finish()
		return nil
	}
}

func RestyTracingBefore(tracer opentracing.Tracer) func(*resty.Client, *resty.Request) error {
	return func(_ *resty.Client, r *resty.Request) error {
		name := strings.Split(r.URL, "?")[0]
		var spanOpts []opentracing.StartSpanOption
		span := opentracing.SpanFromContext(r.Context())
		if span != nil {
			spanOpts = append(spanOpts, opentracing.ChildOf(span.Context()))
		}
		spanOpts = append(spanOpts, []opentracing.StartSpanOption{
			opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			ext.SpanKindRPCClient,
		}...)
		newSpan, newCtx := opentracing.StartSpanFromContextWithTracer(r.Context(), tracer, name, spanOpts...)
		_ = tracer.Inject(newSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if jaegerContext, ok := newSpan.Context().(jaeger.SpanContext); ok {
			var traceID, spanID string
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
			newCtx = context.WithValue(newCtx, kopentracing.TraceKey{}, traceID)
			newCtx = context.WithValue(newCtx, kopentracing.SpanKey{}, spanID)
		}
		r.SetContext(newCtx)
		return nil
	}
}
