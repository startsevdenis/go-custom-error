package CustomErrors

import (
	"fmt"
)

type ErrorExternal struct {
	Code    int
	Resp    string
	Message string
}

func NewErrorExternal(code int, resp string, message string) *ErrorExternal {
	return &ErrorExternal{Code: code, Resp: resp, Message: message}
}

func (e *ErrorExternal) Error() string {
	return fmt.Sprintf("External call error: code %d, response: ", e.Code, e.Resp)
}

func (e *ErrorExternal) PublicMessage() string {
	return e.Message
}
