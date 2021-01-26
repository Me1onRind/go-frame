package file_service

import (
	"context"
	"go-frame/global"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"time"

	"path/filepath"

	minio "github.com/minio/minio-go"
)

type minioService struct {
	getObjectOption *minio.GetObjectOptions
	putObjectOption *minio.PutObjectOptions
	minioClient     *minio.Client
}

func newMinioService() *minioService {
	m := &minioService{
		getObjectOption: &minio.GetObjectOptions{},
		putObjectOption: &minio.PutObjectOptions{},
		minioClient:     global.MinioClient,
	}
	return m
}

func (m *minioService) Upload(ctx *custom_ctx.Context, localFilePath string, timeout time.Duration) (fileIndex string, err *errcode.Error) {
	var c context.Context = ctx
	if timeout > 0 {
		var cancel context.CancelFunc
		c, cancel = context.WithTimeout(c, timeout)
		defer cancel()
	}

	objectName := filepath.Base(localFilePath)
	_, e := m.minioClient.FPutObjectWithContext(c, global.MinioSetting.BucketName, objectName, localFilePath, *m.putObjectOption)
	if e != nil {
		return "", errcode.MinioError.WithError(e)
	}

	return objectName, nil
}

func (m *minioService) Download(ctx *custom_ctx.Context, fileIndex, localFilePath string, timeout time.Duration) (err *errcode.Error) {
	objectName := fileIndex
	var c context.Context = ctx
	if timeout > 0 {
		var cancel context.CancelFunc
		c, cancel = context.WithTimeout(c, timeout)
		defer cancel()
	}

	e := m.minioClient.FGetObjectWithContext(c, global.MinioSetting.BucketName, objectName, localFilePath, *m.getObjectOption)
	if e != nil {
		return errcode.MinioError.WithError(e)
	}

	return nil
}
