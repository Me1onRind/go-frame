package audio

import (
	"go-frame/global"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"

	"go.uber.org/zap"
)

type AudioDao struct{}

func NewAudioDao() *AudioDao {
	a := &AudioDao{}
	return a
}

func (a *AudioDao) Create(ctx *custom_ctx.Context, audio *Audio) *errcode.Error {
	db := ctx.WriteDB(global.DefaultDB)
	if err := db.Save(audio).Error; err != nil {
		ctx.Logger().Error("Create audio record failed", zap.Any("audio", audio))
		return errcode.DBError.WithError(err)
	}
	return nil
}

func (a *AudioDao) UpdateFileStatus(ctx *custom_ctx.Context, id uint64, status uint8) *errcode.Error {
	db := ctx.WriteDB(global.DefaultDB)
	if err := db.Update("file_status", status).Where("id=?", id).Error; err != nil {
		ctx.Logger().Error("Create audio record failed", zap.Uint64("id", id), zap.Uint8("status", status), zap.Error(err))
		return errcode.DBError.WithError(err)
	}
	return nil
}

func (a *AudioDao) GetUnsyncAudioList(ctx *custom_ctx.Context, minID uint64, limit int) ([]*Audio, *errcode.Error) {
	var result []*Audio
	db := ctx.ReadDB(global.DefaultDB)
	if err := db.Where("id > ?", minID).Limit(limit).Find(&result).Error; err != nil {
		ctx.Logger().Error("Get unsync audio failed", zap.Uint64("minID", minID), zap.Error(err))
		return nil, errcode.DBError.WithError(err)
	}
	return result, nil
}
