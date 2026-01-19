package CustomErrors

import (
	"fmt"
)

type ErrorExternal struct {
	Code    int
	Resp    string
	Message string
	Err     error
}

func NewErrorExternal(code int, resp string, message string, err error) *ErrorExternal {
	return &ErrorExternal{Code: code, Resp: resp, Message: message, Err: err}
}

func (e *ErrorExternal) Error() string {
	return fmt.Sprintf("External call error: code %d, response: %s", e.Code, e.Resp)
}

func (e *ErrorExternal) PublicMessage() string {
	return e.Message
}

func (e *ErrorExternal) Unwrap() error {
	return e.Err
}
