package protocol

type ListResp struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}

type ListReq struct {
	Page     int `json:"page" form:"page" binding:"page"`
	PageSize int `json:"page_size" form:"page_size" binding:"page_size"`
}
