package instance

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/stepanovdmitrii/go-esent/internal/pkg/esent"
)

// Instance Instance of the database engine for use in a single process
type Instance struct {
	instancePointer uintptr
}

// CreateInstance Creates a new instance of the database engine
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetcreateinstance-function
func CreateInstance(name string) (*Instance, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("instance name must be specified")
	}

	bytePtr, cErr := syscall.BytePtrFromString(name)
	if cErr != nil {
		return nil, fmt.Errorf("failed to create byte pointer from name: %w", cErr)
	}

	var instancePtr uintptr

	if err := esent.Syscall(esent.JetCreateInstance, uintptr(unsafe.Pointer(&instancePtr)), uintptr(unsafe.Pointer(bytePtr))); err != nil {
		return nil, err
	}

	return &Instance{instancePointer: instancePtr}, nil
}

// Init Puts the database engine into a state where it can support application use of database files
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetinit-function
func (i *Instance) Init() error {
	return esent.Syscall(esent.JetInit, uintptr(unsafe.Pointer(&i.instancePointer)))
}

// SetSystemParameter Set configuration setting of the database engine
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetsetsystemparameter-function
//func (i *Instance) SetSystemParameter(parameter *config.SystemParameter) error {
//	if parameter == nil {
//		return fmt.Errorf("parameter must be specified")
//	}
//	return dll.Syscall(dll.JetSetSystemParameter, uintptr(unsafe.Pointer(&i.instancePointer)), 0, uintptr(82), uintptr(2147483648), uintptr(0))
//}
