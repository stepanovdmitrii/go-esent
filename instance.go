package goesent

import (
	"fmt"
	"syscall"
	"unsafe"
)

// Instance Instance of the database engine for use in a single process
type Instance struct {
	instancePointer uintptr
}

// CreateInstance Creates a new instance of the database engine
// Name must be unique within process
func CreateInstance(name string) (*Instance, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("instance name must be specified")
	}

	bytePtr, cErr := syscall.BytePtrFromString(name)
	if cErr != nil {
		return nil, fmt.Errorf("failed to create byte pointer from name: %w", cErr)
	}

	var instancePtr uintptr

	_, _, err := syscall.SyscallN(esentCreateInstance, uintptr(unsafe.Pointer(&instancePtr)), uintptr(unsafe.Pointer(bytePtr)))
	if err != 0 && err < 0 {
		return nil, err
	}

	return &Instance{instancePointer: instancePtr}, nil
}

func (i *Instance) Init() error {
	_, _, err := syscall.SyscallN(esentInit, uintptr(unsafe.Pointer(&i.instancePointer)))
	if err != 0 && err < 0 {
		return err
	}
	return nil
}

// SetSystemParameter Set configuration parameter
//func (i *Instance) SetSystemParameter(sysParameter SystemParameter) error {
//TODO
//}
