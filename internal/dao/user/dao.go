package user

import (
	"go-frame/global"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"gorm.io/gorm"
)

type UserDao struct {
	dbKey string
}

func NewUserDao() *UserDao {
	return &UserDao{
		dbKey: global.DefaultDB,
	}
}

func (u *UserDao) GetUserByUserID(ctx *custom_ctx.Context, userID uint64) (*User, *errcode.Error) {
	var user User
	err := ctx.ReadDB(u.dbKey).WithContext(ctx).Where("user_id=?", userID).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errcode.RecordNotFound
	}
	if err != nil {
		return nil, errcode.DBError.WithError(err)
	}
	return &user, nil
}

func (u *UserDao) GetUserByUsername(ctx *custom_ctx.Context, username string) (*User, *errcode.Error) {
	var user User
	err := ctx.ReadDB(u.dbKey).Where("username = ?", username).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errcode.DBError.WithError(err)
	}
	return &user, nil
}
