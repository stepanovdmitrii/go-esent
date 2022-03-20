package config

import "fmt"

// Transaction Log Parameters
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/transaction-log-parameters

// SysParameterBaseName JET_paramBaseName
func SysParameterBaseName(value string) (*SystemParameter, error) {
	if len(value) != 3 {
		return nil, fmt.Errorf("baseName value must be 3 characters long")
	}
	return sysParameterString(3, value)
}

// SysParameterCircularLog JET_paramCircularLog
func SysParameterCircularLog(value bool) (*SystemParameter, error) {
	return sysParameterBool(17, value)
}

// SysParameterLogFilePath JET_paramLogFilePath
func SysParameterLogFilePath(value string) (*SystemParameter, error) {
	return sysParameterString(2, value)
}

// SysParameterSystemPath JET_paramSystemPath
func SysParameterSystemPath(value string) (*SystemParameter, error) {
	return sysParameterString(0, value)
}
