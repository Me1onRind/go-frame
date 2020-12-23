package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/utils/session"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := session.GetUserInfo(c)
		if userInfo == nil {
			c.Abort()
			return
		}

		c.Next()
	}
}
