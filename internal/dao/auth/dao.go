package auth

import (
	"go-frame/global"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"gorm.io/gorm"
)

type AuthDao struct {
}

func NewAuthDao() *AuthDao {
	return &AuthDao{}
}

func (a *AuthDao) ListAuths(ctx *context.Context, page int, pageSize int) ([]*Auth, int64, *errcode.Error) {
	var list []*Auth
	var total int64
	db := ctx.ReadDB(global.DefaultDB)
	if err := db.Limit(pageSize).Offset((page - 1) * pageSize).Preload("Config").Find(&list).Error; err != nil {
		return nil, 0, errcode.DBError.WithError(err)
	}

	if err := db.Model(list).Count(&total).Error; err != nil {
		return nil, 0, errcode.DBError.WithError(err)
	}

	return list, total, nil
}

func (a *AuthDao) GetAuthByAppKey(ctx *context.Context, appKey string) (*Auth, *errcode.Error) {
	var auth Auth
	if err := ctx.ReadDB(global.DefaultDB).Where("app_key=?", appKey).Preload("Config").Take(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errcode.DBError.WithError(err)
	}

	return &auth, nil
}
