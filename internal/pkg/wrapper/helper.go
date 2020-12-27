package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/metadata"
	uuid "github.com/satori/go.uuid"
	"go-frame/global"
)

func getRequestID(ctx context.Context) string {
	requestID, _ := metadata.Get(ctx, global.ProtocolRequestIDKey)
	if len(requestID) == 0 {
		requestID = uuid.NewV4().String()
	}

	return requestID
}

func getJWTToken(ctx context.Context) string {
	jwtToken, _ := metadata.Get(ctx, global.ProtocolJWTTokenKey)
	return jwtToken
}
