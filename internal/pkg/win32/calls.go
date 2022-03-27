package win32

import (
	"fmt"
	"syscall"
)

// StringToBytePosinter ...
func StringToBytePosinter(str string) (*byte, error) {
	bytePtr, cErr := syscall.BytePtrFromString(str)
	if cErr != nil {
		return nil, fmt.Errorf("failed to create byte pointer from string '%s': %w", str, cErr)
	}
	return bytePtr, nil
}
