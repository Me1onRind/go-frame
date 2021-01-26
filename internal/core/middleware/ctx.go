package middleware

import (
	"go-frame/global"
	"go-frame/internal/core/custom_ctx"

	"github.com/gin-gonic/gin"
)

func InitContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := custom_ctx.NewContext(global.Logger, nil)
		custom_ctx.LoadIntoGinContext(ctx, c)
	}
}
