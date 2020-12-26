package errcode

import (
	"fmt"
	"go-frame/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	OriginError error  `json:"-"`
}

var usedCodes = map[int]struct{}{}

func NewError(code int, msg string) *Error {
	if _, ok := usedCodes[code]; ok {
		panic(fmt.Sprintf("Error code:%d is duplicate", code))
	}
	usedCodes[code] = struct{}{}
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

func toGrpcCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case ServerErrorCode:
		statusCode = codes.Internal
	case InvalidParamCode:
		statusCode = codes.InvalidArgument
	case JWTAuthorizedFailCode, JWTTimeoutCode:
		statusCode = codes.Unauthenticated
	case RecordNotFoundCode:
		statusCode = codes.NotFound
	default:
		statusCode = codes.Unknown
	}
	return statusCode
}

func (e *Error) ToRpcError() error {
	grpcCode := toGrpcCode(e.Code)
	s, _ := status.New(grpcCode, e.Msg).WithDetails(&pb.Error{Errcode: int32(e.Code), Message: e.Msg})
	return s.Err()
}
