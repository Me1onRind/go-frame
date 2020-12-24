package routers

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/controller/http/user"
	"go-frame/internal/pkg/gateway"
	"go-frame/internal/protocol/user_proto"
)

func registerUserApi(router *gin.RouterGroup) {
	userController := user.NewUserContoller()

	router.POST("/login", gateway.Json(userController.Login, &user_proto.LoginReq{}))
}
