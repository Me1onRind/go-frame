package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	newContext "go-frame/internal/pkg/context"
	"go-frame/internal/pkg/logger"
	"time"
)

func Logger(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		newCtx := newContext.GetCommonContext(ctx)
		logger.WithTrace(newCtx).WithFields(
			logger.KV("fullMethod", req.Method()),
			logger.KV("req", req.Body()),
		).Info("GRPC Request begin")

		startTime := time.Now()
		err := fn(ctx, req, resp)
		end := time.Now()

		logger.WithTrace(newCtx).WithFields(
			logger.KV("fullMethod", req.Method()),
			logger.KV("req", req.Body()),
			logger.KV("resp", resp),
			logger.KV("error", err),
			logger.KV("cost", end.Sub(startTime)),
		).Info("GRPC Request completed")
		return err
	}
}
