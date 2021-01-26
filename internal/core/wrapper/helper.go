package wrapper

import (
	"context"
	"go-frame/internal/constant/proto_constant"

	"github.com/micro/go-micro/v2/metadata"
)

func getJWTToken(ctx context.Context) string {
	jwtToken, _ := metadata.Get(ctx, proto_constant.ProtocolJWTTokenKey)
	return jwtToken
}
