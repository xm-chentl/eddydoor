package common

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return c.Message
}

func ResponseErr(code int, format string, args ...interface{}) error {
	return &CustomError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
