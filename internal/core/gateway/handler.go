package gateway

import (
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"reflect"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler func(c *custom_ctx.Context, raw interface{}) (data interface{}, err *errcode.Error)

func Json(handler Handler, paramType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := custom_ctx.GetFromGinContext(c)
		var raw interface{}
		if paramType != nil {
			requestParams := parserProtocol(paramType)
			if err := c.ShouldBind(requestParams); err != nil {
				c.JSON(200, errcode.InvalidParam.WithError(err))
				return
			}
			raw = requestParams
		}

		var data interface{}
		var e *errcode.Error
		defer func() {
			ctx.Logger().Info("Json gateway", zap.Reflect("reqObj", raw), zap.Reflect("respObj", raw))
		}()
		data, e = handler(ctx, raw)
		if e == nil {
			e = errcode.Success
		}

		span := ctx.Span()
		span.SetTag("errcode", e.Code)
		span.SetTag("message", e.Msg)

		c.JSON(200, NewResponse(e, data))
	}
}

func parserProtocol(paramType interface{}) interface{} {
	if paramType == nil {
		return nil
	}
	valueType := reflect.TypeOf(paramType)
	kind := valueType.Kind()
	if kind == reflect.Ptr {
		return reflect.New(valueType.Elem()).Interface()
	}
	if kind == reflect.Struct {
		return reflect.New(valueType).Interface()
	}
	return nil
}
