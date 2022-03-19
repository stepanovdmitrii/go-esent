package goesent

import (
	"fmt"
	"syscall"
)

// EsentError Esent engine error
type EsentError struct {
	Code    uintptr
	Message string
}

// Error ...
func (e *EsentError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func handleEsentErr(err syscall.Errno) error {
	if err >= 0 {
		return nil
	}

}
