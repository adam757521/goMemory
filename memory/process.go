package memory

import (
	"strings"
	"unsafe"
)

type Process uintptr

type ProcessEntry struct {
	DwSize    uint32
	Padding   [4]byte
	ProcessId uint64
	Padding1  [28]byte
	Name      [256]byte
}

func (snapshot Process) ProcessFirst(process *ProcessEntry) error {
	_, _, err := callProc(kernelProcessFirst, uintptr(snapshot), uintptr(unsafe.Pointer(process)))
	return err
}

func (snapshot Process) ProcessNext(process *ProcessEntry) error {
	_, _, err := callProc(kernelProcessNext, uintptr(snapshot), uintptr(unsafe.Pointer(process)))
	return err
}

func OpenProcess(desiredAccess int, inheritHandle bool, processId uint32) (Process, error) {
	inheritHandleInt := int8(0)
	if inheritHandle {
		inheritHandleInt = 1
	}

	handle, _, err := callProc(kernelOpenProcess, uintptr(desiredAccess), uintptr(inheritHandleInt), uintptr(processId))
	return Process(handle), err
}

func (p Process) CloseProcess() error {
	_, _, err := callProc(kernelCloseProcess, uintptr(p))
	return err
}

func FindProcess(name string) (*ProcessEntry, error) {
	snapshot, err := CreateToolSnapshot(0x00000002, 0)
	if err != nil {
		return nil, err
	}

	process := ProcessEntry{}
	process.DwSize = uint32(unsafe.Sizeof(process))

	err = snapshot.ProcessFirst(&process)
	if err != nil {
		println("Could not get first process.")
		_ = snapshot.CloseProcess()
		return nil, err
	}

	for {
		processName := string(process.Name[:])

		if strings.Index(processName, name) == 0 {
			_ = snapshot.CloseProcess()
			return &process, nil
		}

		err = snapshot.ProcessNext(&process)
		if err != nil {
			break
		}
	}

	_ = snapshot.CloseProcess()
	return nil, err
}

func (p Process) GetProcessId() (uint32, error) {
	processId, _, err := callProc(kernelGetProcessId, uintptr(p))
	return uint32(processId), err
}
