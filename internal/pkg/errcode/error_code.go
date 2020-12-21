package errcode

const (
	SuccessCode         int = 0
	ServerErrorCode     int = 1000000
	InvalidParamCode    int = 1000001
	DBErrorCode         int = 1000002
	RedisErrorCode      int = 1000003
	RecordNotFoundCode  int = 1000004
	CopyStructErrorCode int = 1000005
)

var (
	Success         = NewError(SuccessCode, "Success")
	ServerError     = NewError(ServerErrorCode, "Server Internal Error")
	InvalidParam    = NewError(InvalidParamCode, "Invalid Param")
	DBError         = NewError(DBErrorCode, "DB Error")
	RedisError      = NewError(RedisErrorCode, "Redis Error")
	RecordNotFound  = NewError(RecordNotFoundCode, "Record Not Found")
	CopyStructError = NewError(CopyStructErrorCode, "Copy Struct Error")
)
