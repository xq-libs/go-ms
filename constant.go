package ms

const (
	SuccessCode      = 0
	ServerErrorCode  = 1001
	UnknownErrorCode = 1002

	DBErrorCode         = 1101
	DBNotFoundErrorCode = 1102
)

var (
	Success      = NewMessage(SuccessCode, "ms.success", "Success")
	ServerError  = NewMessage(ServerErrorCode, "ms.server.error", "Server error")
	UnknownError = NewMessage(UnknownErrorCode, "ms.server.error.unknown", "Unknown error")

	DBError         = NewMessage(DBErrorCode, "ms.server.error.db.error", "Db error")
	DBNotFoundError = NewMessage(DBNotFoundErrorCode, "ms.server.error.db.notFound", "Not Found Record")
)
