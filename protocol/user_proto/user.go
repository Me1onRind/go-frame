package user_proto

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserInfoReq struct {
	UserID uint64 `form:"user_id" binding:"required"`
}

type GetUserInfoByTokenReq struct {
	Token string `form:"token" json:"token" binding:"required"`
}
