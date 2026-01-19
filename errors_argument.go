package CustomErrors

import (
	"fmt"
	"strings"
)

type ErrorMissedArgument struct {
	Message         string
	MissedArguments []string
}

func (e *ErrorMissedArgument) PublicMessage() string {
	return internal_error
}

func (e *ErrorMissedArgument) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, strings.Join(e.MissedArguments, ", "))
}

func NewErrorMissedArgument(missedArguments ...string) *ErrorMissedArgument {
	return &ErrorMissedArgument{
		Message:         argument_error,
		MissedArguments: missedArguments,
	}
}
