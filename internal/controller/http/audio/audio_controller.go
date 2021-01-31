package audio

import (
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"go-frame/internal/service/audio"
	"go-frame/protocol"
	"go-frame/protocol/audio_proto"
	"os"
)

type AudioController struct {
	audioService *audio.AudioService
}

func NewAudioController() *AudioController {
	a := &AudioController{
		audioService: audio.NewAudioService(),
	}
	return a
}

func (a *AudioController) UploadAudio(ctx *custom_ctx.Context, raw interface{}) (interface{}, *errcode.Error) {
	fileHeader, err := ctx.GinCtx.FormFile("file")
	if err != nil {
		return nil, errcode.UploadFileError.WithError(err)
	}

	localFilepath := "/tmp/" + fileHeader.Filename
	err = ctx.GinCtx.SaveUploadedFile(fileHeader, localFilepath)
	if err != nil {
		return nil, errcode.FileOperationError.WithError(err)
	}
	defer os.Remove(localFilepath)

	e := a.audioService.UploadAudio(ctx, fileHeader.Filename, localFilepath)
	if e != nil {
		return nil, e
	}

	return nil, nil
}

func (a *AudioController) List(ctx *custom_ctx.Context, raw interface{}) (interface{}, *errcode.Error) {
	params := raw.(*audio_proto.ListAudioReq)
	result, total, err := a.audioService.Search(ctx, params.Page, params.PageSize)
	if err != nil {
		return nil, err
	}

	return &protocol.ListResp{
		Page:     params.Page,
		PageSize: params.PageSize,
		Total:    total,
		List:     result,
	}, nil
}
