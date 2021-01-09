package user

import (
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"go-frame/internal/dao/user"
	"go-frame/internal/lib/remote_service/user_service"
	"go-frame/proto/pb"
)

type UserService struct {
	UserDao           *user.UserDao
	RemoteUserService *user_service.RemoteUserService
}

func NewUserService() *UserService {
	return &UserService{
		UserDao:           user.NewUserDao(),
		RemoteUserService: user_service.NewRemoteUserService(),
	}
}

func (u *UserService) Login(ctx context.Context, username string, password string) (*user.User, *errcode.Error) {
	user, err := u.UserDao.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errcode.InvalidLoginParamError
	}

	return user, nil
}

func (u *UserService) GetUserByUserID(ctx context.Context, userID uint64, cache bool) (*user.User, *errcode.Error) {
	return nil, errcode.ServerError
	userInfo, err := u.UserDao.GetUserByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (u *UserService) GetUserFromRemote(ctx context.Context, userID uint64) (*pb.UserInfo, *errcode.Error) {
	return u.RemoteUserService.GetUserInfoByUserID(ctx, userID)
}
