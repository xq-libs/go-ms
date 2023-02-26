package ms

const (
	SuccessCode      = 0
	ServerErrorCode  = 1001
	UnknownErrorCode = 1002

	DBErrorCode         = 1101
	DBNotFoundErrorCode = 1102

	RequestQueryErrorCode     = 1201
	RequestParamErrorCode     = 1202
	RequestBodyBindErrorCode  = 1203
	RequestQueryBindErrorCode = 1204
)

var (
	Success      = NewMessage(SuccessCode, "MSSuccess", "Success")
	ServerError  = NewMessage(ServerErrorCode, "MSError", "Server error")
	UnknownError = NewMessage(UnknownErrorCode, "MSErrorUnknown", "Unknown error")

	DBError         = NewMessage(DBErrorCode, "MSErrorDBError", "Db error")
	DBNotFoundError = NewMessage(DBNotFoundErrorCode, "MSErrorDBNotFound", "Not Found Record")

	RequestQueryError     = NewMessage(RequestQueryErrorCode, "MSErrorQueryRequired", "Request query param {{.param}} missed.")
	RequestParamError     = NewMessage(RequestParamErrorCode, "MSErrorParamRequired", "Request path param {{.param}} missed.")
	RequestBodyBindError  = NewMessage(RequestBodyBindErrorCode, "MSErrorBodyBindError", "Request body bind error.")
	RequestQueryBindError = NewMessage(RequestQueryBindErrorCode, "MSErrorQueryBindError", "Request query param bind error.")
)
