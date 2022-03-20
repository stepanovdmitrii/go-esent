package esent

import "syscall"

var (
	esentDll = syscall.MustLoadDLL("esent.dll")

	JetCreateInstance, _      = syscall.GetProcAddress(esentDll.Handle, "JetCreateInstance")
	JetInit, _                = syscall.GetProcAddress(esentDll.Handle, "JetInit")
	JetStopServiceInstance, _ = syscall.GetProcAddress(esentDll.Handle, "JetStopServiceInstance")
	JetTerm, _                = syscall.GetProcAddress(esentDll.Handle, "JetTerm")
	JetSetSystemParameter, _  = syscall.GetProcAddress(esentDll.Handle, "JetSetSystemParameter")
)
