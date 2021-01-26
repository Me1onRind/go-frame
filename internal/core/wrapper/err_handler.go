package wrapper

import (
	"context"
	customCtx "go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"

	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/server"
	"go.uber.org/zap"
)

func ErrHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		err := fn(ctx, req, resp)
		newCtx := customCtx.GetFromContext(ctx)
		span := newCtx.Span()
		if err != nil {
			if e, ok := err.(*errors.Error); ok {
				span.SetTag("errcode", e.Code)
				span.SetTag("message", e.Detail)
				if e.Code == 500 {
					newCtx.Logger().Error("Internal Server Error", zap.Error(err))
				}
			}
		} else {
			span.SetTag("errcode", errcode.SuccessCode)
			span.SetTag("message", errcode.Success.Msg)
		}

		return err
	}
}
