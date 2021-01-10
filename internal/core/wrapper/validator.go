package wrapper

import (
	"context"
	"go-frame/global"
	"go-frame/internal/core/errcode"

	"github.com/micro/go-micro/v2/server"
)

func Validator(fn server.HandlerFunc) server.HandlerFunc {
	return func(c context.Context, req server.Request, resp interface{}) error {
		if err := global.Validate.Struct(req.Body()); err != nil {
			return errcode.InvalidParam.WithError(err).ToRpcError()
		}
		return fn(c, req, resp)
	}
}
