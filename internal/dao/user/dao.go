package user

import (
	"go-frame/global"
	"go-frame/internal/model"
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
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

func (u *UserDao) GetUserByUserID(ctx context.Context, userID uint64) (*model.User, *errcode.Error) {
	var user model.User
	err := ctx.ReadDB(u.dbKey).Where("user_id = ?", userID).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errcode.DBError.WithError(err)
	}
	return &user, nil
}

func (u *UserDao) UpdateUser(ctx context.Context, user *model.User) *errcode.Error {
	if err := ctx.WriteDB(u.dbKey).Save(user).Error; err != nil {
		return errcode.DBError.WithError(err)
	}
	return nil
}

//func (u *UserDao) Save(ctx context.Context, user *model.User) *errcode.Error {
//return ctx.ReadDB(global.DefaultDB).Save(&user).Error
//}
