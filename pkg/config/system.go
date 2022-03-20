package config

import (
	"fmt"
	"syscall"
	"unsafe"
)

// SystemParameter System parameter
type SystemParameter struct {
	Id        uintptr
	IntPtr    uintptr
	StringPtr uintptr
}

func sysParameterString(id uintptr, str string) (*SystemParameter, error) {
	ptr, err := syscall.BytePtrFromString(str)
	if err != nil {
		return nil, fmt.Errorf("failed to create byte array from string: %w", err)
	}
	return &SystemParameter{
		Id:        id,
		StringPtr: uintptr(unsafe.Pointer(ptr)),
	}, nil
}

func sysParameterBool(id uintptr, value bool) (*SystemParameter, error) {
	var valuePtr uintptr
	if value {
		valuePtr = 1
	}
	return &SystemParameter{
		Id:     id,
		IntPtr: valuePtr,
	}, nil
}

func sysParameterUIntPtr(id uintptr, value uintptr) (*SystemParameter, error) {
	return &SystemParameter{
		Id:     id,
		IntPtr: value,
	}, nil
}
