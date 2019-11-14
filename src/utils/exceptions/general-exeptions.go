package exceptions

import "fmt"

type ConfigurationError struct {
	cause string
}

func (e *ConfigurationError) Error() string {
	return fmt.Sprintf("%s - %s","ConfigurationError", e.cause)
}

func NewConfigurationError(cause string) *ConfigurationError {
	return &ConfigurationError{cause: cause}
}

func CheckAndPanic(e error) {
	if e != nil {
		panic(e)
	}
}
