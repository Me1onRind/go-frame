package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/global"
	"go-frame/internal/lib/auth"
	"go-frame/internal/lib/session"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/gateway"
	"go-frame/internal/utils/ctx_helper"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := session.GetUserInfo(c)
		if userInfo == nil {
			c.JSON(200, gateway.NewResponse(errcode.UnLoginError, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(global.ProtocolJWTTokenKey)
		ctx := ctx_helper.GetHttpContext(c)
		if err := auth.JWTAuth(ctx, token); err != nil {
			c.JSON(200, gateway.NewResponse(err, nil))
			c.Abort()
		}

		c.Next()
	}
}
