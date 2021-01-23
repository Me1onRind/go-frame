package middleware

import (
	"go-frame/global"
	customCtx "go-frame/internal/core/context"

	"github.com/gin-gonic/gin"
)

func InitContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := customCtx.NewContext(global.Logger, nil)
		customCtx.LoadIntoGinContext(ctx, c)
	}
}
