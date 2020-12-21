package user

import (
	"github.com/jinzhu/copier"
	"go-frame/internal/dao/user"
	"go-frame/internal/model"
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

func (u *UserService) UpdateUser(ctx context.Context, updateInfo *UpdateInfo) (*model.User, *errcode.Error) {
	user, err := u.UserDao.GetUserByUserID(ctx, updateInfo.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errcode.RecordNotFound.WithInfof("UserID[%d] not exist", updateInfo.UserID)
	}

	if err := copier.Copy(&user, &updateInfo); err != nil {
		return nil, errcode.CopyStructError.WithError(err)
	}

	if err := u.UserDao.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
