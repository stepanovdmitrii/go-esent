//go:build windows && amd64

package esent

import (
	"fmt"
	"unsafe"

	"github.com/stepanovdmitrii/go-esent/internal/pkg/esent"
	"github.com/stepanovdmitrii/go-esent/internal/pkg/win32"
	"github.com/stepanovdmitrii/go-esent/pkg/config"
)

// Instance Instance of the database engine for use in a single process
type Instance struct {
	handle uintptr
}

// CreateInstance Creates a new instance of the database engine
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetcreateinstance-function
func CreateInstance(name string) (*Instance, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("instance name must be specified")
	}

	bytePtr, cErr := win32.StringToBytePosinter(name)
	if cErr != nil {
		return nil, cErr
	}

	var handle uintptr

	if err := esent.Syscall(esent.JetCreateInstance, uintptr(unsafe.Pointer(&handle)), uintptr(unsafe.Pointer(bytePtr))); err != nil {
		return nil, err
	}

	return &Instance{handle: handle}, nil
}

// Init Puts the database engine into a state where it can support application use of database files
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetinit-function
func (i *Instance) Init() error {
	if i == nil {
		return fmt.Errorf("instance is not initialized")
	}
	return esent.Syscall(esent.JetInit, uintptr(unsafe.Pointer(&i.handle)))
}

// Term Initiates the shutdown of an instance
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetterm-function
func (i *Instance) Term() error {
	if i == nil {
		return fmt.Errorf("instance is not initialized")
	}
	return esent.Syscall(esent.JetTerm, i.handle)
}

// SetSystemParameter Set configuration setting of the database engine
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetsetsystemparameter-function
func (i *Instance) SetSystemParameter(parameter *config.SystemParameter) error {
	if i == nil {
		return fmt.Errorf("instance is not initialized")
	}
	if parameter == nil {
		return fmt.Errorf("parameter must be specified")
	}
	return esent.Syscall(esent.JetSetSystemParameter, uintptr(unsafe.Pointer(&i.handle)), 0, parameter.Id, parameter.IntPtr, parameter.StringPtr)
}

// BeginSession Starts a session
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetbeginsession-function
func (i *Instance) BeginSession() (*Session, error) {
	if i == nil {
		return nil, fmt.Errorf("instance is not initialized")
	}
	var handle uintptr
	if err := esent.Syscall(esent.JetBeginSession, i.handle, uintptr(unsafe.Pointer(&handle))); err != nil {
		return nil, err
	}

	return &Session{handle: handle}, nil
}
