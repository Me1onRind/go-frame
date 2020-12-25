package routers

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/controller/http/jwt"
	"go-frame/internal/pkg/gateway"
	"go-frame/protocol/jwt_proto"
)

func registerJWTApi(router *gin.RouterGroup) {
	jwtController := jwt.NewJWTController()

	router = router.Group("/jwt")
	router.GET("generate_token", gateway.Json(jwtController.GenerateToken, &jwt_proto.GenerateTokenReq{}))
}
