package audio

import (
	"go-frame/global"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"

	"go.uber.org/zap"
)

type AudioDao struct{}

func NewAudioDao() *AudioDao {
	a := &AudioDao{}
	return a
}

func (a *AudioDao) Create(ctx *context.Context, audio *Audio) *errcode.Error {
	db := ctx.WriteDB(global.DefaultDB)
	if err := db.Save(audio).Error; err != nil {
		ctx.Logger().Error("Create audio record failed", zap.Any("audio", audio))
		return errcode.DBError.WithError(err)
	}
	return nil
}

func (a *AudioDao) UpdateFileStatus(ctx *context.Context, id uint64, status uint8) *errcode.Error {
	db := ctx.WriteDB(global.DefaultDB)
	if err := db.Update("file_status", status).Where("id=?", id).Error; err != nil {
		ctx.Logger().Error("Create audio record failed", zap.Uint64("id", id), zap.Uint8("status", status))
		return errcode.DBError.WithError(err)
	}
	return nil
}
