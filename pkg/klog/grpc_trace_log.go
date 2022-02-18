package klog

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// ClientInterceptor grpc client wrapper
func GRPCUnaryClientTraceLog(l *Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		nctx := WithTraceCtx(l, ctx)
		return invoker(nctx, method, req, reply, cc, opts...)
	}
}

func GRPCUnaryServerTraceLog(l *Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		nctx := WithTraceCtx(l, ctx)
		return handler(nctx, req)
	}
}

func GRPCUnaryServerRequestLogging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		beginTime := time.Now()
		logger := FromTraceCtx(ctx)
		reply, err := handler(ctx, req)
		latency := time.Since(beginTime).Microseconds()
		if err != nil {
			logger.Error().Str("method", info.FullMethod).Int64("latency(μs)", latency).Interface("req", req).Interface("reply", reply).Err(err).Msg("grpc_server_request")
			return reply, err
		}
		logger.Info().Str("method", info.FullMethod).Int64("latency(μs)", latency).Interface("req", req).Interface("reply", reply).Msg("grpc_server_request")
		return reply, err
	}
}

func GRPCUnaryClientRequestLogging() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		beginTime := time.Now()
		logger := FromTraceCtx(ctx)
		err := invoker(ctx, method, req, reply, cc, opts...)
		latency := time.Since(beginTime).Microseconds()
		if err != nil {
			logger.Error().Str("method", method).Int64("latency(μs)", latency).Interface("req", req).Interface("reply", reply).Err(err).Msg("grpc_server_request")
			return err
		}
		logger.Info().Str("method", method).Int64("latency(μs)", latency).Interface("req", req).Interface("reply", reply).Msg("grpc_client_request")
		return err
	}
}
