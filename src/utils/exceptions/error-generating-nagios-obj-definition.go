package exceptions

import "fmt"

type ErrorGeneratingNagiosObjDefinition struct {
	cause string
}

func (e *ErrorGeneratingNagiosObjDefinition) Error() string {
	return fmt.Sprintf("%s - %s","ErrorGeneratingNagiosObjDefinition", e.cause)
}

func NewErrorGeneratingNagiosObjDefinition(cause string) *ErrorGeneratingNagiosObjDefinition {
	return &ErrorGeneratingNagiosObjDefinition{cause: cause}
}