package gateway

import (
	"go-frame/internal/pkg/errcode"
)

type Response struct {
	Errcode int         `json:"errcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(err *errcode.Error, data interface{}) *Response {
	if data == nil {
		data = struct{}{}
	}
	return &Response{
		Errcode: err.Code,
		Message: err.Msg,
		Data:    data,
	}
}
