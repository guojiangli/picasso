package kjaeger

import (
	"context"

	"github.com/guojiangli/picasso/pkg/trace/kopentracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func HTTPServerTracing(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
		parentSpanContext, _ := tracer.Extract(opentracing.HTTPHeaders, carrier)
		spanOpts := []opentracing.StartSpanOption{
			opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			ext.SpanKindRPCServer,
			opentracing.Tag{Key: string(ext.HTTPMethod), Value: c.Request.Method},
			ext.RPCServerOption(parentSpanContext),
		}
		span := tracer.StartSpan(c.Request.Host+c.Request.RequestURI, spanOpts...)
		defer span.Finish()
		if jaegerContext, ok := span.Context().(jaeger.SpanContext); ok {
			var traceID, spanID string
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
			c.Request = c.Request.WithContext(
				opentracing.ContextWithSpan(
					context.WithValue(
						context.WithValue(c.Request.Context(), kopentracing.TraceKey{}, traceID),
						kopentracing.SpanKey{},
						spanID,
					),
					span,
				),
			)
		} else {
			c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), span))
		}
		c.Next()
		ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))
	}
}
