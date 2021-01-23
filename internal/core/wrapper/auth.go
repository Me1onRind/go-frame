package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	customCtx "go-frame/internal/core/context"
	"go-frame/internal/lib/auth"
)

func JWT(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		token := getJWTToken(ctx)
		newCtx := customCtx.GetFromContext(ctx)
		if err := auth.JWTAuth(newCtx, token); err != nil {
			return err.ToRpcError()
		}

		return fn(ctx, req, resp)
	}
}
