package config

import "fmt"

// Temporary Database Parameters
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/temporary-database-parameters

// SysParameterTempPath JET_paramTempPath
func SysParameterTempPath(value string) (*SystemParameter, error) {
	if len(value) > 247 {
		return nil, fmt.Errorf("tempPath must be a maximum of 247 characters")
	}
	return sysParameterString(1, value)
}
