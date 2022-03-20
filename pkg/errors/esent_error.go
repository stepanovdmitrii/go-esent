package errors

import (
	"fmt"
)

// EsentError Esent engine error
type EsentError struct {
	Code    ErrorCode
	Message string
}

// Error ...
func (e *EsentError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
