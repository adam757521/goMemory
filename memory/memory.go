package memory

import "unsafe"

func (p Process) ReadMemoryPointer(address uintptr, read unsafe.Pointer, size uintptr) error {
	_, _, err := callProc(kernelReadProcess, uintptr(p), address, uintptr(read), size, Nil)
	return err
}

func (p Process) ReadMemory(address uintptr) (uintptr, error) {
	var read uintptr
	read = uintptr(unsafe.Pointer(&read))

	err := p.ReadMemoryPointer(address, unsafe.Pointer(&read), unsafe.Sizeof(read))
	return read, err
}

func (p Process) WriteMemoryPointer(address uintptr, pointer unsafe.Pointer, size uintptr) error {
	_, _, err := callProc(kernelWriteProcess, uintptr(p), address, uintptr(pointer), size, Nil)
	return err
}

func (p Process) WriteIntMemory(address uintptr, value int) error {
	// To support -1 writes
	return p.WriteMemoryPointer(address, unsafe.Pointer(&value), unsafe.Sizeof(value))
}

func (p Process) WriteMemory(address uintptr, data uintptr) error {
	return p.WriteMemoryPointer(address, unsafe.Pointer(&data), unsafe.Sizeof(data))
}

func (p Process) FindPattern(pattern string, module *Module) (uintptr, error) {
	return 0, nil
}
