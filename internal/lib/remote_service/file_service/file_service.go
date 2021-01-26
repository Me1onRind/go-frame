package file_service

import (
	"fmt"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"time"
)

type FileStoreType uint

const (
	Minio FileStoreType = iota
)

type FileService interface {
	Upload(ctx *custom_ctx.Context, localFilePath string, timeout time.Duration) (fileIndex string, err *errcode.Error)
	Download(ctx *custom_ctx.Context, fileIndex, localFilePath string, timeout time.Duration) (err *errcode.Error)
}

func NewFileService(fileStoreType FileStoreType) FileService {
	switch fileStoreType {
	case Minio:
		return newMinioService()
	default:
		panic(fmt.Sprintf("Unsupport fileStoreType:%d", fileStoreType))
	}
}
