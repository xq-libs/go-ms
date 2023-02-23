package ms

type Error struct {
	Cause   error
	Message Message
}

func NewError(err error, msg Message) Error {
	return Error{
		Cause:   err,
		Message: msg,
	}
}

func (e Error) Code() int64 {
	return e.Message.Code
}

func (e Error) Error() string {
	return e.Message.DefaultValue
}
