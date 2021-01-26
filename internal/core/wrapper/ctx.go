package wrapper

import (
	"context"
	"go-frame/global"
	customCtx "go-frame/internal/core/custom_ctx"

	"github.com/micro/go-micro/v2/server"
)

func InitContext(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		c := customCtx.NewContext(global.Logger, ctx)
		ctx = customCtx.LoadIntoContext(c, ctx)
		return fn(ctx, req, resp)
	}
}
