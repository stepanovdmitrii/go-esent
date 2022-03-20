package instance

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/stepanovdmitrii/go-esent/internal/pkg/dll"
	"github.com/stepanovdmitrii/go-esent/internal/pkg/errors"
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

	_, _, err := syscall.SyscallN(dll.JetCreateInstance, uintptr(unsafe.Pointer(&instancePtr)), uintptr(unsafe.Pointer(bytePtr)))
	if esentErr := errors.HandleEsentErr(err); esentErr != nil {
		return nil, esentErr
	}

	return &Instance{instancePointer: instancePtr}, nil
}

// Init Puts the database engine into a state where it can support application use of database files
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetinit-function
func (i *Instance) Init() error {
	_, _, err := syscall.SyscallN(dll.JetInit, uintptr(unsafe.Pointer(&i.instancePointer)))
	if esentErr := errors.HandleEsentErr(err); esentErr != nil {
		return esentErr
	}
	return nil
}
