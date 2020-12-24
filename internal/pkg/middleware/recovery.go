package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/pkg/logger"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				tracer := getTracer(c)
				logger.WithTrace(tracer).Errorf("Panic recover err:%v, stack:\n%s", err, debug.Stack())
				c.String(500, "Server Internal Error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
