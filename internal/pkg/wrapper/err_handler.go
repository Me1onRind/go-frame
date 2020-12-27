package wrapper

import (
	"context"
	"github.com/micro/go-micro/errors"
	//"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	newContext "go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/logger"
)

func ErrHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		err := fn(ctx, req, resp)
		if err != nil {
			if e, ok := err.(*errors.Error); ok {
				if e.Code == 500 {
					newCtx := newContext.GetCommonContext(ctx)
					logger.WithTrace(newCtx).Errorf("Internal Server Error:%s", e.Detail)
					err = errcode.ServerError.ToRpcError()
				}
			}
		}

		return err
	}
}
