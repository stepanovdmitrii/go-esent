package dll

import "syscall"

var (
	esentDll = syscall.MustLoadDLL("esent.dll")

	JetCreateInstance, _ = syscall.GetProcAddress(esentDll.Handle, "JetCreateInstance")
	JetInit, _           = syscall.GetProcAddress(esentDll.Handle, "JetInit")
)
