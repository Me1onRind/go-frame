package user_service

import (
	"github.com/micro/go-micro/v2/client"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"go-frame/internal/lib/client/grpc"
	"go-frame/proto/pb"
	//"time"
)

const (
	jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiNGZhOGJiNGI2Mjc1ZTg4MDFkZjk4NjI1OWE1YjU3YTAiLCJhcHBfc2VjcmV0IjoiMmY3NTJiZGNlMzYzZDVjZDZjZWE2OWU1ODFhY2Q1MDYifQ.PAORHbZVumzQU6IUbY-d4l6CK8KBjQLgW7AyVm5Vs4E"
)

type RemoteUserService struct {
	UserRpcClient client.Client
}

func NewRemoteUserService() *RemoteUserService {
	return &RemoteUserService{
		UserRpcClient: grpc.GoFrameClient,
	}
}

func (r *RemoteUserService) GetUserInfoByUserID(ctx context.Context, userID uint64) (*pb.UserInfo, *errcode.Error) {
	req := r.UserRpcClient.NewRequest("go-frame-grpc", "UserService.GetUserInfo", &pb.GetUserReq{UserID: userID})
	userInfo := &pb.UserInfo{}
	if err := r.UserRpcClient.Call(grpc.JWTContext(ctx, jwtToken), req, userInfo); err != nil {
		return nil, errcode.FromRpcError(err)
	}

	return userInfo, nil
}
