package user

import (
	"github.com/jinzhu/copier"
	"go-frame/internal/service/user"
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
)

type UserController struct {
	UserService *user.UserService
}

func NewUserContoller() *UserController {
	return &UserController{
		UserService: user.NewUserService(),
	}
}

func (u *UserController) GetUserByID(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*GetUserInfoRequest)
	return map[string]interface{}{
		"userID":   request.UserID,
		"username": "un",
	}, nil
}

func (u *UserController) SearchUsers(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*SearchUsersRequest)
	return map[string]interface{}{
		"list": []map[string]interface{}{
			{
				"userID":   1,
				"userName": "user1",
				"userType": request.UserType,
			},
		},
	}, nil
}

func (u *UserController) UpdateUser(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*UpdateUserRequest)

	var updateInfo user.UpdateInfo
	if err := copier.Copy(&updateInfo, request); err != nil {
		return nil, errcode.CopyStructError.WithError(err)
	}

	user, err := u.UserService.UpdateUser(ctx, &updateInfo)
	if err != nil {
		return nil, err
	}
	return user, nil
}
