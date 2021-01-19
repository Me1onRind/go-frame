package file_service

import (
	"fmt"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"time"
)

type FileStoreType uint

const (
	Minio FileStoreType = iota
)

type FileService interface {
	Upload(ctx context.Context, localFilePath string, timeout time.Duration) (fileIndex string, err *errcode.Error)
	Download(ctx context.Context, fileIndex, localFilePath string, timeout time.Duration) (err *errcode.Error)
}

func NewFileService(fileStoreType FileStoreType) FileService {
	switch fileStoreType {
	case Minio:
		return newMinioService()
	default:
		panic(fmt.Sprintf("Unsupport fileStoreType:%d", fileStoreType))
	}
}
