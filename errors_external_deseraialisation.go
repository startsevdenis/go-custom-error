package CustomErrors

import "fmt"

type ErrorExternalDeserialisation struct {
	Content string
	Message string
}

func NewErrorExternalDeserialisation(content string, message string) *ErrorExternalDeserialisation {
	return &ErrorExternalDeserialisation{Content: content, Message: message}
}

func (e *ErrorExternalDeserialisation) Error() string {
	return fmt.Sprintf("Error External Deserialisation: %s, %s", e.Message, e.Content)
}

func (e *ErrorExternalDeserialisation) PublicMessage() string {
	return internal_error
}
