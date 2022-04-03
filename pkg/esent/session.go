package esent

import (
	"fmt"
	"unsafe"

	"github.com/stepanovdmitrii/go-esent/internal/pkg/esent"
	"github.com/stepanovdmitrii/go-esent/internal/pkg/win32"
)

// Session ...
type Session struct {
	sesID uintptr
}

// CreateDatabase Creates and attaches a database file to be used with the ESE database engine
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetcreatedatabase-function
func (s *Session) CreateDatabase(filename string, flags CreateDatabaseFlag) (*Database, error) {
	if s == nil {
		return nil, fmt.Errorf("session is not initialized")
	}
	bytePtr, cErr := win32.StringToBytePosinter(filename)
	if cErr != nil {
		return nil, cErr
	}
	var handle uintptr
	if err := esent.Syscall(esent.JetCreateDatabase, s.sesID, uintptr(unsafe.Pointer(bytePtr)), Null, uintptr(unsafe.Pointer(&handle)), uintptr(flags)); err != nil {
		return nil, err
	}
	return &Database{
		sesID: s.sesID,
		dbID:  handle,
	}, nil
}

// EndSession Ends the session, and cleans up and deallocates any resources associated with the specified session.
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetendsession-function
func (s *Session) EndSession() error {
	if s == nil {
		return fmt.Errorf("session is not initialized")
	}
	return esent.Syscall(esent.JetEndSession, s.sesID)
}
