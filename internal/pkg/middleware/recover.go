package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/gateway"
	"go-frame/internal/pkg/logger"
	"runtime/debug"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				tracer := getTracer(c)
				logger.WithTrace(tracer).Errorf("Panic recover err:%v, stack:\n%s", err, debug.Stack())
				c.JSON(200, gateway.NewResponse(errcode.ServerError, nil))
				c.Abort()
			}
		}()
		c.Next()
	}
}
