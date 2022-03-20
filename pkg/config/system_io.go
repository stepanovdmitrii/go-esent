package config

// I/O Parameters
// https://docs.microsoft.com/en-us/windows/win32/extensible-storage-engine/i-o-parameters

// SysParameterCreatePathIfNotExist JET_paramCreatePathIfNotExist
func SysParameterCreatePathIfNotExist(value bool) (*SystemParameter, error) {
	return sysParameterBool(100, value)
}
