package user

import (
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"go-frame/internal/dao/auth"
	"go-frame/internal/dao/user"
	"go-frame/internal/lib/remote_service/user_service"
	"go-frame/proto/pb"

	"github.com/google/uuid"
)

type UserService struct {
	UserDao           *user.UserDao
	RemoteUserService *user_service.RemoteUserService
	AutoDao           *auth.AuthDao
}

func NewUserService() *UserService {
	return &UserService{
		UserDao:           user.NewUserDao(),
		RemoteUserService: user_service.NewRemoteUserService(),
	}
}

func (u *UserService) Login(ctx *custom_ctx.Context, username string, password string) (string, *errcode.Error) {
	usr, err := u.UserDao.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if u == nil {
		return "", errcode.RecordNotFound
	}

	if usr.Password != password {
		return "", errcode.InvalidLoginParamError
	}

	cacheUserInfo := &user.CacheUserInfo{
		UserID:    usr.UserID,
		Username:  usr.Username,
		RoleNames: []string{},
	}

	if roles, err := u.AutoDao.GetRolesByRoleIDs(ctx, usr.RoleIDs); err != nil {
		return "", err
	} else {
		for _, r := range roles {
			cacheUserInfo.RoleNames = append(cacheUserInfo.RoleNames, r.RoleName)
		}
	}

	token, e := uuid.NewRandom()
	if e != nil {
		return "", errcode.ServerError.WithError(e)
	}

	strToken := token.String()
	if err := u.UserDao.StroeLoginInfo(ctx, cacheUserInfo, strToken); err != nil {
		return "", err
	}

	return strToken, nil
}

func (u *UserService) GetUserInfoByToken(ctx *custom_ctx.Context, token string) (*user.CacheUserInfo, *errcode.Error) {
	return u.UserDao.GetUserByToken(ctx, token)
}

func (u *UserService) GetUserByUserID(ctx *custom_ctx.Context, userID uint64, cache bool) (*user.User, *errcode.Error) {
	userInfo, err := u.UserDao.GetUserByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (u *UserService) GetUserFromRemote(ctx *custom_ctx.Context, userID uint64) (*pb.UserInfo, *errcode.Error) {
	return u.RemoteUserService.GetUserInfoByUserID(ctx, userID)
}
