package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/global"
	"go-frame/internal/lib/auth"
	"go-frame/internal/pkg/gateway"
	"go-frame/internal/utils/session"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := session.GetUserInfo(c)
		if userInfo == nil {
			c.Abort()
			return
		}

		c.Next()
	}
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(global.ProtocolJWTTokenKey)
		if err := auth.JWTAuth(getTracer(c), token); err != nil {
			c.JSON(200, gateway.NewResponse(err, nil))
			c.Abort()
		}

		c.Next()
	}
}
