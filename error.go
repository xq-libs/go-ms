package ms

import "fmt"

type Error struct {
	Cause   error
	Message Message
}

func NewError(msg Message) Error {
	return Error{
		Cause:   nil,
		Message: msg,
	}
}

func NewError2(err error, msg Message) Error {
	return Error{
		Cause:   err,
		Message: msg,
	}
}

func (e Error) Code() int64 {
	return e.Message.Code
}

func (e Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Message: %s, error: %v", e.Message.GetDefaultMessage(), e.Cause)
	}
	return e.Message.GetDefaultMessage()
}
