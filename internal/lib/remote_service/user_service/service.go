package user_service

import (
	"go-frame/internal/lib/client/grpc"
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/proto/pb"
	//"time"
)

const (
	jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiNGZhOGJiNGI2Mjc1ZTg4MDFkZjk4NjI1OWE1YjU3YTAiLCJhcHBfc2VjcmV0IjoiMmY3NTJiZGNlMzYzZDVjZDZjZWE2OWU1ODFhY2Q1MDYifQ.PAORHbZVumzQU6IUbY-d4l6CK8KBjQLgW7AyVm5Vs4E"
)

type RemoteUserService struct {
	RpcUserService pb.UserService
}

func NewRemoteUserService() *RemoteUserService {
	return &RemoteUserService{
		RpcUserService: grpc.GoFrameClient,
	}
}

func (r *RemoteUserService) GetUserInfoByUserID(ctx context.Context, userID uint64) (*pb.UserInfo, *errcode.Error) {
	userInfo, err := r.RpcUserService.GetUserInfo(grpc.JwtContext(ctx, jwtToken), &pb.GetUserReq{
		UserID: userID,
	})
	if err != nil {
		return nil, errcode.FromRpcError(err)
	}
	return userInfo, nil
}
