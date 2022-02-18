package klog

import (
	"context"

	"picasso/pkg/trace/kopentracing"
	"github.com/rs/zerolog"
)

func WithTraceCtx(l *Logger, ctx context.Context) context.Context {
	newLogger := l.Logger.With().Logger()
	lctx := newLogger.With()
	if traceKey, ok := ctx.Value(kopentracing.TraceKey{}).(string); ok {
		lctx = lctx.Str("trace_key", traceKey)
	}
	if spanKey, ok := ctx.Value(kopentracing.SpanKey{}).(string); ok {
		lctx = lctx.Str("span_key", spanKey)
	}
	newLogger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return lctx
	})
	return newLogger.WithContext(ctx)
}

func FromTraceCtx(ctx context.Context) *Logger {
	logger := zerolog.Ctx(ctx)
	if level := logger.GetLevel(); level == zerolog.Disabled {
		newLogger := defaultLogger.With().Logger()
		return NewLoggerWith(&newLogger)
	}
	return &Logger{
		Logger: logger,
	}
}
