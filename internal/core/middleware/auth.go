package middleware

import (
	"go-frame/internal/constant/proto_constant"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"go-frame/internal/core/gateway"
	"go-frame/internal/core/session"
	"go-frame/internal/lib/auth"

	"github.com/gin-gonic/gin"
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
		token := c.GetHeader(proto_constant.ProtocolJWTTokenKey)
		ctx := custom_ctx.GetFromGinContext(c)
		if err := auth.JWTAuth(ctx, token); err != nil {
			c.JSON(200, gateway.NewResponse(err, nil))
			c.Abort()
		}

		c.Next()
	}
}
