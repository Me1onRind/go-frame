package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/utils/ctx_helper"
	"time"
)

func AccessLogger(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		newCtx := ctx_helper.GetCommonContext(ctx)
		logger.WithTrace(newCtx).WithFields(
			logger.KV("fullMethod", req.Method()),
			logger.KV("req", req.Body()),
		).Info("GRPC Request Start")

		startTime := time.Now()
		err := fn(ctx, req, resp)
		end := time.Now()

		logger.WithTrace(newCtx).WithFields(
			logger.KV("fullMethod", req.Method()),
			logger.KV("req", req.Body()),
			logger.KV("resp", resp),
			logger.KV("error", err),
			logger.KV("cost", end.Sub(startTime)),
		).Info("GRPC Request End")
		return err
	}
}
