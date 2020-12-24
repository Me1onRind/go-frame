package user

import (
	"go-frame/internal/dao/user"
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
)

type UserService struct {
	UserDao *user.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		UserDao: user.NewUserDao(),
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
