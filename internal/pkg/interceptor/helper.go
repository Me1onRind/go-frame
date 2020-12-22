package interceptor

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go-frame/global"
	"go-frame/internal/pkg/logger"
	"google.golang.org/grpc/metadata"
)

func setRequestID(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	values := md.Get(global.ProtocolRequestIDKey)

	var xRequestID string
	if len(values) == 0 {
		xRequestID = uuid.NewV4().String()
	} else {
		xRequestID = values[0]
	}

	return context.WithValue(ctx, global.ContextRequestIDKey, xRequestID)
}

func getTracer(c context.Context) *logger.SimpleTracer {
	requestID := c.Value(global.ContextRequestIDKey).(string)
	return &logger.SimpleTracer{
		ReqID: requestID,
	}
}
