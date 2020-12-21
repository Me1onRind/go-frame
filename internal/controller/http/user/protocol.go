package user

type GetUserInfoRequest struct {
	UserID string `form:"user_id" binding:"required"`
}

type SearchUsersRequest struct {
	UserType uint8 `form:"user_type" binding:"userType"`
}

type UpdateUserRequest struct {
	UserID   uint64  `json:"user_id" binding:"required"`
	UserType *uint8  `json:"user_type" binding:"omitempty,userType"`
	GroupID  *uint64 `json:"group_id"`
}
