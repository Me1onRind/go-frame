package gateway

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/utils/ctx_helper"
	"reflect"
)

type Handler func(c *context.HttpContext) (data interface{}, err *errcode.Error)

func Json(handler Handler, paramType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpContext := ctx_helper.GetHttpContext(c)
		if paramType != nil {
			requestParams := parserProtocol(paramType)
			if err := c.ShouldBind(requestParams); err != nil {
				c.JSON(200, errcode.InvalidParam.WithError(err))
				return
			}
			httpContext.Raw = requestParams
		}
		data, e := handler(httpContext)
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
