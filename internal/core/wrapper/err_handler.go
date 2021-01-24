package wrapper

import (
	"context"
	customCtx "go-frame/internal/core/context"

	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/zap"
)

func ErrHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		err := fn(ctx, req, resp)
		if err != nil {
			if e, ok := err.(*errors.Error); ok {
				if e.Code == 500 {
					newCtx := customCtx.GetFromContext(ctx)
					newCtx.Logger().Error("Internal Server Error", zap.Error(err))
				}
			}
		}

		return err
	}
}
