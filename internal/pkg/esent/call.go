package esent

import (
	"syscall"

	"github.com/stepanovdmitrii/go-esent/pkg/errors"
)

var uint32Max int64 = 4294967295

const error_ENVVAR_NOT_FOUND = 203

func Syscall(trap uintptr, args ...uintptr) error {
	r1, _, err := syscall.SyscallN(trap, args...)
	if err != 0 && err != error_ENVVAR_NOT_FOUND {
		return err
	}

	if r1 != 0 {
		var esentErrCode int64 = int64(r1) - uint32Max
		if esentErrCode >= 0 { //warning
			return nil
		}
		code := errors.ErrorCode(esentErrCode)
		if message, ok := errorCodeToMessage[code]; ok {
			return &errors.EsentError{
				Code:    code,
				Message: message,
			}
		}
		return &errors.EsentError{
			Code:    code,
			Message: "Unknown error",
		}
	}
	return nil
}
