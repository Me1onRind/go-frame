package audio

import (
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"go-frame/internal/dao/audio"
	"go-frame/internal/lib/remote_service/file_service"
	"time"
)

type AudioService struct {
	fileService file_service.FileService
	audioDao    *audio.AudioDao
}

func NewAudioService() *AudioService {
	a := &AudioService{
		fileService: file_service.NewFileService(file_service.Minio),
		audioDao:    audio.NewAudioDao(),
	}
	return a
}

func (a *AudioService) UploadAudio(ctx *custom_ctx.Context, filename, localFilepath string) *errcode.Error {
	existFile, err := a.audioDao.GetAudioByFilename(ctx, filename)
	if err != nil {
		return err
	}

	if existFile != nil {
		return errcode.RecordExistError
	}

	fileIndex, err := a.fileService.Upload(ctx, localFilepath, time.Second*5)
	if err != nil {
		return err
	}

	audioEntity := &audio.Audio{
		Filename:   filename,
		StoreIndex: fileIndex,
	}
	if err := a.audioDao.Create(ctx, audioEntity); err != nil {
		return err
	}

	return nil
}

func (a *AudioService) Search(ctx *custom_ctx.Context, page, pageSize int) (result []*audio.Audio, total int64, err *errcode.Error) {
	result, total, err = a.audioDao.Search(ctx, map[string]interface{}{}, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (a *AudioService) CleanUnsync(ctx *custom_ctx.Context, args []interface{}) *errcode.Error {
	id := args[0].(uint64)
	_, err := a.audioDao.GetAudioByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
