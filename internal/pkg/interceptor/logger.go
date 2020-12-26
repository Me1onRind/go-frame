package interceptor

import (
	"context"
	newContext "go-frame/internal/pkg/context"
	"go-frame/internal/pkg/logger"
	"google.golang.org/grpc"
	"time"
)

func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	newCtx := ctx.(*newContext.CommonContext)
	logger.WithTrace(newCtx).WithFields(
		logger.KV("fullMethod", info.FullMethod),
		logger.KV("req", req),
		logger.KV("clientIP", info.FullMethod),
	).Info("Request begin")

	startTime := time.Now()
	resp, err := handler(ctx, req)
	end := time.Now()

	logger.WithTrace(newCtx).WithFields(
		logger.KV("fullMethod", info.FullMethod),
		logger.KV("req", req),
		logger.KV("resp", resp),
		logger.KV("cost", end.Sub(startTime)),
	).Info("Request completed")
	return resp, err
}
