package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	newContext "go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/logger"
	"runtime/debug"
)

func Recover(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) (err error) {
		defer func() {
			newCtx := ctx.(*newContext.CommonContext)
			if e := recover(); e != nil {
				logger.WithTrace(newCtx).Errorf("Panic recover err:%v, stack:\n%s", e, debug.Stack())
				err = errcode.ServerError.ToRpcError()
			}
		}()
		err = fn(ctx, req, resp)
		return
	}
}
