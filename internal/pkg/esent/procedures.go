package esent

import "syscall"

var (
	esentDll = syscall.MustLoadDLL("esent.dll")

	JetCreateInstance      = mustGetProcAddress("JetCreateInstance")
	JetInit                = mustGetProcAddress("JetInit")
	JetStopServiceInstance = mustGetProcAddress("JetStopServiceInstance")
	JetTerm                = mustGetProcAddress("JetTerm")
	JetSetSystemParameter  = mustGetProcAddress("JetSetSystemParameter")
	JetBeginSession        = mustGetProcAddress("JetBeginSession")
	JetCreateDatabase      = mustGetProcAddress("JetCreateDatabase")
	JetEndSession          = mustGetProcAddress("JetEndSession")
	JetCloseDatabase       = mustGetProcAddress("JetCloseDatabase")
)

func mustGetProcAddress(proc string) uintptr {
	handle, err := syscall.GetProcAddress(esentDll.Handle, proc)
	if err != nil {
		panic(err)
	}
	return handle
}
