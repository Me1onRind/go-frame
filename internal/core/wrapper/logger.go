package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	"go-frame/internal/utils/ctx_helper"
	"go.uber.org/zap"
	"time"
)

func AccessLogger(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		newCtx := ctx_helper.GetCommonContext(ctx)
		newCtx.Logger().Info("GRPC Request Start",
			zap.String("method", req.Method()),
			zap.Any("reqBody", req.Body()),
		)

		startTime := time.Now()
		err := fn(ctx, req, resp)
		end := time.Now()

		newCtx.Logger().Info("GRPC Request Start",
			zap.String("method", req.Method()),
			zap.Any("reqBody", req.Body()),
			zap.Any("resp", resp),
			zap.Error(err),
			zap.Duration("cost", end.Sub(startTime)),
		)

		return err
	}
}
