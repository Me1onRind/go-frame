package user

import (
	"encoding/json"
	"go-frame/global"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"strconv"

	"gorm.io/gorm"
)

const (
	userTokenPrefix = "user:token:"
)

type UserDao struct {
	dbKey string
}

func NewUserDao() *UserDao {
	return &UserDao{
		dbKey: global.DefaultDB,
	}
}

func (u *UserDao) GetUserByUserID(ctx *custom_ctx.Context, userID uint64) (*User, *errcode.Error) {
	var user User
	err := ctx.ReadDB(u.dbKey).WithContext(ctx).Where("user_id=?", userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errcode.RecordNotFound
	}
	if err != nil {
		return nil, errcode.DBError.WithError(err)
	}
	return &user, nil
}

func (u *UserDao) GetUserByUsername(ctx *custom_ctx.Context, username string) (*User, *errcode.Error) {
	var user User
	err := ctx.ReadDB(u.dbKey).Where("username = ?", username).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errcode.DBError.WithError(err)
	}
	return &user, nil
}

func (u *UserDao) StroeLoginInfo(ctx *custom_ctx.Context, user *CacheUserInfo, token string) *errcode.Error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return errcode.JsonMarchalError.WithError(err)
	}
	_, err = global.Redis.Eval(ctx,
		"local exit_token = redis.call('hget', 'login:users', KEYS[3])\n"+
			"if exit_token then redis.call('del', KEYS[1] .. exit_token) end\n"+
			"local set_result = redis.call('set', KEYS[1] .. KEYS[2], ARGV[1], 'ex', ARGV[2])\n"+
			"return redis.call('hset', 'login:users', KEYS[3], ARGV[3])",
		[]string{userTokenPrefix, token, strconv.FormatUint(user.UserID, 10)},
		userJson, 3600*2, token).Result()

	if err != nil {
		return errcode.RedisError.WithError(err)
	}

	return nil
}

func (u *UserDao) GetUserByToken(ctx *custom_ctx.Context, token string) (*CacheUserInfo, *errcode.Error) {
	result, err := global.Redis.Get(ctx, userTokenPrefix+token).Bytes()
	if err != nil {
		return nil, errcode.RedisError.WithError(err)
	}

	var user CacheUserInfo
	if err := json.Unmarshal(result, &user); err != nil {
		return nil, errcode.JsonUnMarchalError.WithError(err)
	}
	return &user, nil
}
