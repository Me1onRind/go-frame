package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"go-frame/internal/core/gateway"
	"runtime/debug"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx := context.GetFromGinContext(c)
				ctx.Logger().Sugar().Errorf("Panic recover err:%v, stack:\n%s", err, debug.Stack())
				c.JSON(200, gateway.NewResponse(errcode.ServerError, nil))
				c.Abort()
			}
		}()
		c.Next()
	}
}
