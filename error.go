package ms

import (
	"errors"
	"gorm.io/gorm"
)

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

func WrapDbError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Error{
			Cause:   err,
			Message: DBNotFoundError,
		}
	} else if err != nil {
		return Error{
			Cause:   err,
			Message: DBError,
		}
	} else {
		return nil
	}
}
