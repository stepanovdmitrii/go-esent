package goesent

import "syscall"

var (
	esentDll = syscall.MustLoadDLL("esent.dll")

	esentCreateInstance, _ = syscall.GetProcAddress(esentDll.Handle, "JetCreateInstance")
	esentInit, _           = syscall.GetProcAddress(esentDll.Handle, "JetInit")
)

type esent struct {
	dll *syscall.DLL
}
