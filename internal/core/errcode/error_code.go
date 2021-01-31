package errcode

const (
	SuccessCode             int = 0
	ServerErrorCode         int = 1000000
	InvalidParamCode        int = 1000001
	DBErrorCode             int = 1000002
	RedisErrorCode          int = 1000003
	RecordNotFoundErrorCode int = 1000004
	CopyStructErrorCode     int = 1000005
	RemoteServerErrorCode   int = 1000006
	MinioErrorCode          int = 1000007
	SaveSessionErrorCode    int = 1000008
	OptExtractErrorCode     int = 1000009
	UploadFileErrorcode     int = 1000010
	FileOperationErrorcode  int = 1000011
	RecordExistErrorCode    int = 1000012

	InvalidLoginParamCode int = 2000000
	UnLoginCode           int = 2000001

	JWTSignedFailCode     int = 2000100
	JWTAuthorizedFailCode int = 2000101
	JWTTimeoutCode        int = 2000102
	AppSecretWrongCode    int = 2000103
)

var (
	Success            = NewError(SuccessCode, "Success")
	ServerError        = NewError(ServerErrorCode, "Server Internal Error")
	InvalidParam       = NewError(InvalidParamCode, "Invalid Param")
	DBError            = NewError(DBErrorCode, "DB Error")
	RedisError         = NewError(RedisErrorCode, "Redis Error")
	RecordNotFound     = NewError(RecordNotFoundErrorCode, "Record Not Found")
	CopyStructError    = NewError(CopyStructErrorCode, "Copy Struct Error")
	RemoteServerError  = NewError(RemoteServerErrorCode, "Remote Server Error")
	MinioError         = NewError(MinioErrorCode, "Minio Error")
	SaveSessionError   = NewError(SaveSessionErrorCode, "Save Session Error")
	UploadFileError    = NewError(UploadFileErrorcode, "Upload File Error")
	FileOperationError = NewError(FileOperationErrorcode, "File Operation Error")
	RecordExistError   = NewError(RecordExistErrorCode, "Record Exist")

	InvalidLoginParamError = NewError(InvalidLoginParamCode, "Invalid Username Or Password")
	UnLoginError           = NewError(UnLoginCode, "Not Login")

	JWTSignedFailError     = NewError(JWTSignedFailCode, "JWT Signed Fail")
	JWTAuthorizedFailError = NewError(JWTAuthorizedFailCode, "JWT Authorized Fail")
	JWTTimeoutError        = NewError(JWTTimeoutCode, "JWT Timeout")
	AppSecretWrongError    = NewError(AppSecretWrongCode, "App Secret Wrong")
	OptExtractError        = NewError(OptExtractErrorCode, "Decode Opentracing Carrier Fail")
)
