package esent

import (
	"fmt"

	"github.com/stepanovdmitrii/go-esent/internal/pkg/esent"
)

// Database ...
type Database struct {
	sessionHandle  uintptr
	databaseHandle uintptr
}

// CloseDatabase Closes a database file
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetclosedatabase-function
func (db *Database) CloseDatabase() error {
	if db == nil {
		return fmt.Errorf("database is not initialized")
	}
	return esent.Syscall(esent.JetCloseDatabase, db.sessionHandle, db.databaseHandle)
}
