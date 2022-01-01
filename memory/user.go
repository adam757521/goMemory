package memory

func GetForegroundWindow() (Process, error) {
	handle, _, err := callProc(userGetForegroundWindow)
	return Process(handle), err
}

func GetKeyState(key int) int16 {
	state, _, _ := callProc(userGetKeyState, uintptr(key))
	return int16(state)
}
