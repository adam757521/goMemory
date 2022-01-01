package memory

func CreateToolSnapshot(dwFlags int, processID uint32) (Process, error) {
	handle, _, err := callProc(kernelHelpCreateTool, uintptr(dwFlags), uintptr(processID))
	return Process(handle), err
}
