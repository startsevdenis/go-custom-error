package CustomErrors

type ErrorServer struct {
	Message       string
	MessagePublic string
	Err           error
}

func (e *ErrorServer) Error() string {
	return e.Message
}

func (e *ErrorServer) Unwrap() error {
	return e.Err
}

func (e *ErrorServer) PublicMessage() string {
	return e.MessagePublic
}

func NewErrorServer(msg string, err error) *ErrorServer {
	return &ErrorServer{Message: msg, Err: err, MessagePublic: internal_error}
}
