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
		newCtx := ctx.(*newContext.CommonContext)
		logger.WithTrace(newCtx).WithFields(
			logger.KV("fullMethod", req.Method()),
			logger.KV("req", req.Body()),
		).Info("Request begin")

		startTime := time.Now()
		err := fn(ctx, req, resp)
		end := time.Now()

		logger.WithTrace(newCtx).WithFields(
			logger.KV("fullMethod", req.Method()),
			logger.KV("req", req.Body()),
			logger.KV("resp", resp),
			logger.KV("cost", end.Sub(startTime)),
		).Info("Request completed")
		return err
	}
}
