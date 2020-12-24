package user

import (
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/protocol/user_proto"
	"go-frame/internal/service/user"
	"go-frame/internal/utils/session"
)

type UserController struct {
	UserService *user.UserService
}

func NewUserContoller() *UserController {
	return &UserController{
		UserService: user.NewUserService(),
	}
}

func (u *UserController) Login(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*user_proto.LoginReq)
	userInfo, err := u.UserService.Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	sessionUserInfo := &session.UserInfo{
		UserID:   userInfo.ID,
		Username: userInfo.Username,
	}
	session.SetUserInfo(ctx.Context, sessionUserInfo)

	return sessionUserInfo, nil
}
