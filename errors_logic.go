package CustomErrors

type ErrorLogic struct {
	Message string
}

func (el *ErrorLogic) PublicMessage() string {
	return internal_error
}

func NewErrorLogic(message string) *ErrorLogic {
	return &ErrorLogic{Message: message}
}

func (el *ErrorLogic) Error() string {
	return el.Message
}
