package routers

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/controller/http/user"
	"go-frame/internal/pkg/gateway"
)

func registerUserApi(router *gin.RouterGroup, path string) {
	userController := user.NewUserContoller()
	r := router.Group(path)
	r.GET("/get_user_info", gateway.Json(userController.GetUserByID, &user.GetUserInfoRequest{}))
	r.GET("/list", gateway.Json(userController.SearchUsers, &user.SearchUsersRequest{}))

	r.POST("/update", gateway.Json(userController.UpdateUser, &user.UpdateUserRequest{}))
}
