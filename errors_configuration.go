package CustomErrors

import "fmt"

type ErrorConfiguration struct {
	Message   string
	ConfigKey string
	ErrorType string
}

func NewErrorConfiguration(message string, configKey string, errorType string) *ErrorConfiguration {
	return &ErrorConfiguration{Message: message, ConfigKey: configKey, ErrorType: errorType}
}

func (e *ErrorConfiguration) Error() string {
	return fmt.Sprintf("Configuration error: %s %s", e.ErrorType, e.ConfigKey)
}

func (e *ErrorConfiguration) PublicMessage() string {
	return e.Message
}
