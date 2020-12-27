package user_proto

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,passwd"`
}

type GetUserInfoReq struct {
	UserID uint64 `form:"user_id" binding:"required"`
}
