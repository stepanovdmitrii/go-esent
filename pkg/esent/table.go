package esent

import (
	"fmt"

	"github.com/stepanovdmitrii/go-esent/internal/pkg/esent"
)

// Table ...
type Table struct {
	sesID   uintptr
	tableID uintptr
}

// CloseTable Closes an open table
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/jetclosetable-function
func (t *Table) CloseTable() error {
	if t == nil {
		return fmt.Errorf("table is not initialized")
	}
	return esent.Syscall(esent.JetCloseTable, t.sesID, t.tableID)
}
