package errcode

import (
	"fmt"
	"go-frame/proto/pb"
)

type Error struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	OriginError error  `json:"-"`
}

var codes = map[int]struct{}{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("Error code:%d is duplicate", code))
	}
	codes[code] = struct{}{}
	return &Error{code, msg, nil}
}

func (e *Error) WithInfo(info string) *Error {
	newError := &Error{}
	newError.Code = e.Code
	newError.Msg = fmt.Sprintf("%s:%s", e.Msg, info)
	return newError
}

func (e *Error) WithInfof(format string, args ...interface{}) *Error {
	info := fmt.Sprintf(format, args...)
	return e.WithInfo(info)
}

func (e *Error) WithError(err error) *Error {
	newErr := e.WithInfo(err.Error())
	newErr.OriginError = err
	return newErr
}

func (e *Error) ToRpcError() error {
	return &pb.Error{Errcode: int32(e.Code), Message: e.Msg}
}
