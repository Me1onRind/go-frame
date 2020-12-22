package interceptor

import (
	"context"
	"go-frame/internal/pkg/logger"
	"google.golang.org/grpc"
	"time"
)

func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx = setRequestID(ctx)
	logger.WithTrace(getTracer(ctx)).WithFields(
		logger.KV("fullMethod", info.FullMethod),
		logger.KV("req", req),
		logger.KV("clientIP", info.FullMethod),
	).Info("Request begin")

	startTime := time.Now()
	resp, err := handler(ctx, req)
	end := time.Now()

	logger.WithTrace(getTracer(ctx)).WithFields(
		logger.KV("fullMethod", info.FullMethod),
		logger.KV("reqestParams", req),
		logger.KV("resp", resp),
		logger.KV("cost", end.Sub(startTime)),
	).Info("Request completed")
	return resp, err
}
