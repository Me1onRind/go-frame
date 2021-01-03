package routers

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/controller/http/auth"
	"go-frame/internal/core/gateway"
	"go-frame/protocol/auth_proto"
)

func registerJWTApi(router *gin.RouterGroup) {
	authController := auth.NewAuthController()

	router = router.Group("/auth")
	router.GET("list", gateway.Json(authController.ListAuths, nil))
	router.GET("generate_jwt_token", gateway.Json(authController.GenerateToken, &auth_proto.GenerateTokenReq{}))
}
