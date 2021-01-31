package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	customCtx "go-frame/internal/core/custom_ctx"
	"go.uber.org/zap"
	"time"
)

func AccessLog(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		newCtx := customCtx.GetFromContext(ctx)
		var err error

		startTime := time.Now()
		defer func() {
			end := time.Now()
			newCtx.Logger().Info("Access log",
				zap.String("protocol", "grpc"),
				zap.String("method", req.Method()),
				zap.Reflect("reqBody", req.Body()),
				zap.Reflect("resp", resp),
				zap.Error(err),
				zap.Duration("cost", end.Sub(startTime)),
			)
		}()

		err = fn(ctx, req, resp)
		return err
	}
}
