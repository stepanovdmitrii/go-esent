package errors

import (
	"syscall"

	"github.com/stepanovdmitrii/go-esent/pkg/errors"
)

// HandleEsentErr Convert regular syscall.Errno to EsentError
func HandleEsentErr(err syscall.Errno) error {
	var errCode errors.ErrorCode = errors.ErrorCode(err)
	if err >= 0 {
		return nil
	}
	if message, ok := errorCodeToMessage[errCode]; ok {
		return &errors.EsentError{
			Code:    errCode,
			Message: message,
		}
	}
	return &errors.EsentError{
		Code:    errCode,
		Message: "Unknown error",
	}
}
