package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/global"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"go-frame/internal/core/gateway"
	"go-frame/internal/core/session"
	"go-frame/internal/lib/auth"
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
		ctx := context.GetFromGinContext(c)
		if err := auth.JWTAuth(ctx, token); err != nil {
			c.JSON(200, gateway.NewResponse(err, nil))
			c.Abort()
		}

		c.Next()
	}
}
