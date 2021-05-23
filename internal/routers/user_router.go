package routers

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/controller/http/user"
	"go-frame/internal/core/gateway"
	"go-frame/protocol/user_proto"
)

func registerUserApi(router *gin.RouterGroup) {
	userController := user.NewUserContoller()

	router = router.Group("/user")
	router.POST("/login", gateway.Json(userController.Login, &user_proto.LoginReq{}))
	router.GET("/get_user_info", gateway.Json(userController.GetUserInfo, &user_proto.GetUserInfoReq{}))
	router.GET("/info", gateway.Json(userController.GetUserInfoByToken, &user_proto.GetUserInfoByTokenReq{}))
}
