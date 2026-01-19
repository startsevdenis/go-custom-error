package CustomErrors

import "fmt"

type ErrorExternalDeserialisation struct {
	Content string
	Message string
	Err     error
}

func NewErrorExternalDeserialisation(content string, message string, err error) *ErrorExternalDeserialisation {
	return &ErrorExternalDeserialisation{Content: content, Message: message, Err: err}
}

func (e *ErrorExternalDeserialisation) Error() string {
	return fmt.Sprintf("Error External Deserialisation: %s, %s", e.Message, e.Content)
}

func (e *ErrorExternalDeserialisation) PublicMessage() string {
	return internal_error
}

func (e *ErrorExternalDeserialisation) Unwrap() error {
	return e.Err
}
