package esent

import (
	"fmt"
	"unsafe"

	"github.com/stepanovdmitrii/go-esent/internal/pkg/esent"
	"github.com/stepanovdmitrii/go-esent/internal/pkg/win32"
)

// Database ...
type Database struct {
	sesID uintptr
	dbID  uintptr
}

// CloseDatabase Closes a database file
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetclosedatabase-function
func (db *Database) CloseDatabase() error {
	if db == nil {
		return fmt.Errorf("database is not initialized")
	}
	return esent.Syscall(esent.JetCloseDatabase, db.sesID, db.dbID)
}

// CreateTable Creates an empty table
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetcreatetable-function
func (db *Database) CreateTable(name string, pages uint32, density uint32) (*Table, error) {
	if db == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	bytePtr, cErr := win32.StringToBytePosinter(name)
	if cErr != nil {
		return nil, cErr
	}
	var handle uintptr
	if err := esent.Syscall(esent.JetCreateTable, db.sesID, db.dbID, uintptr(unsafe.Pointer(bytePtr)), uintptr(pages), uintptr(density), uintptr(unsafe.Pointer(&handle))); err != nil {
		return nil, err
	}
	return &Table{
		sesID:   db.sesID,
		tableID: handle,
	}, nil
}
