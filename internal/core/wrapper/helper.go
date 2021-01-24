package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/metadata"
	"go-frame/global"
)

func getJWTToken(ctx context.Context) string {
	jwtToken, _ := metadata.Get(ctx, global.ProtocolJWTTokenKey)
	return jwtToken
}
