package CustomErrors

import "fmt"

type ErrorExternalTimeout struct {
	WaitTime int
	Uri      string
	Message  string
}

func (el *ErrorExternalTimeout) PublicMessage() string {
	return request_timeout_error
}

func NewErrorExternalTimeout(uri string, waitTime int) *ErrorExternalTimeout {
	return &ErrorExternalTimeout{
		WaitTime: waitTime,
		Uri:      uri,
		Message:  request_timeout_error,
	}
}

func (el *ErrorExternalTimeout) Error() string {
	return fmt.Sprintf("%s, uri: %s, wait time: %d", el.Message, el.Uri, el.WaitTime)
}
