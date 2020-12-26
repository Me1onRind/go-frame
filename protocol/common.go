package protocol

type ListResp struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}
