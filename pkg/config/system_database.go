package config

import "fmt"

// Database config
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/database-parameters

// DatabasePageSize ...
type DatabasePageSize uintptr

const (
	DatabasePageSize2KB DatabasePageSize = 2048
	DatabasePageSize4KB DatabasePageSize = 4096
	DatabasePageSize8KB DatabasePageSize = 8192
)

// SysParameterDatabasePageSize JET_paramDatabasePageSize
func SysParameterDatabasePageSize(value DatabasePageSize) (*SystemParameter, error) {
	return sysParameterUIntPtr(64, uintptr(value))
}

// SysParameterDbExtensionSize JET_paramDbExtensionSize
func SysParameterDbExtensionSize(value uint32) (*SystemParameter, error) {
	if value < 1 || value > 2147483647 {
		return nil, fmt.Errorf("dbExtensionSize must be in range 1-2147483647")
	}
	return sysParameterUIntPtr(18, uintptr(value))
}
