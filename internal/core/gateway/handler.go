package gateway

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"reflect"
)

type Handler func(c *context.Context, raw interface{}) (data interface{}, err *errcode.Error)

func Json(handler Handler, paramType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.GetFromGinContext(c)
		var raw interface{}
		if paramType != nil {
			requestParams := parserProtocol(paramType)
			if err := c.ShouldBind(requestParams); err != nil {
				c.JSON(200, errcode.InvalidParam.WithError(err))
				return
			}
			raw = requestParams
		}
		data, e := handler(ctx, raw)
		if e == nil {
			e = errcode.Success
		}

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
