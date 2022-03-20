package config

import (
	"fmt"
	"syscall"
	"unsafe"
)

// SystemParameter System parameter
type SystemParameter struct {
	Id        uint
	IntPtr    uintptr
	StringPtr uintptr
}

// SysParameterAlternateDatabaseRecoveryPath
// JET_paramAlternateDatabaseRecoveryPath
func SysParameterAlternateDatabaseRecoveryPath(path string) (*SystemParameter, error) {
	//if path == "" || len(path) > 246 {
	//	return nil, fmt.Errorf("path must be non-empty string max 246 symbols")
	//}
	return sysParameterString(113, path)
}

// SysParameterCleanupMismatchedLogFiles
// JET_paramCleanupMismatchedLogFiles
func SysParameterCleanupMismatchedLogFiles(cleanup bool) (*SystemParameter, error) {
	return sysParameterBool(77, cleanup)
}

// SysParameterDeleteOutOfRangeLogs
// JET_paramDeleteOutOfRangeLogs
func SysParameterDeleteOutOfRangeLogs(delete bool) (*SystemParameter, error) {
	return sysParameterBool(52, delete)
}

// SysParameterOSSnapshotTimeout
// JET_paramOSSnapshotTimeout
func SysParameterOSSnapshotTimeout(milliseconds int32) (*SystemParameter, error) {
	if milliseconds < 0 {
		return nil, fmt.Errorf("milliseconds value must be greater than zero")
	}
	return sysParameterInt32(82, milliseconds)
}

// SysParameterZeroDatabaseDuringBackup
// JET_paramZeroDatabaseDuringBackup
func SysParameterZeroDatabaseDuringBackup(zero bool) (*SystemParameter, error) {
	return sysParameterBool(71, zero)
}

func sysParameterString(id uint, str string) (*SystemParameter, error) {
	ptr, err := syscall.BytePtrFromString(str)
	if err != nil {
		return nil, fmt.Errorf("failed to create byte array from string: %w", err)
	}
	return &SystemParameter{
		Id:        id,
		StringPtr: uintptr(unsafe.Pointer(ptr)),
	}, nil
}

func sysParameterBool(id uint, value bool) (*SystemParameter, error) {
	var valuePtr uintptr
	if value {
		valuePtr = 1
	}
	return &SystemParameter{
		Id:     id,
		IntPtr: valuePtr,
	}, nil
}

func sysParameterInt32(id uint, value int32) (*SystemParameter, error) {
	return &SystemParameter{
		Id:     id,
		IntPtr: uintptr(value),
	}, nil
}
