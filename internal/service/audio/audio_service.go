package audio

import (
	"go-frame/internal/dao/audio"
	"go-frame/internal/lib/remote_service/file_service"
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

func (a *AudioService) CreateAudio() {
}
