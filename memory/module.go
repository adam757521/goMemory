package memory

import (
	"strings"
	"unsafe"
)

type Module struct {
	DwSize   uint32
	Padding  [16]byte
	Address  *byte
	Size     uint32
	Padding1 [12]byte
	Name     [256]byte
	Padding2 [260]byte
}

func (snapshot Process) ModuleFirst(module *Module) error {
	_, _, err := callProc(kernelModuleFirst, uintptr(snapshot), uintptr(unsafe.Pointer(module)))
	return err
}

func (snapshot Process) ModuleNext(module *Module) error {
	_, _, err := callProc(kernelModuleNext, uintptr(snapshot), uintptr(unsafe.Pointer(module)))
	return err
}

func FindBaseModule(pid uint32, closeSnapshot bool) (Process, *Module, error) {
	snapshot, err := CreateToolSnapshot(0x00000010|0x00000008, pid)
	if err != nil {
		return snapshot, nil, err
	}

	module := Module{}
	module.DwSize = uint32(unsafe.Sizeof(module))

	err = snapshot.ModuleFirst(&module)
	if err != nil {
		println("Could not get first module.")
		_ = snapshot.CloseProcess()
		return snapshot, nil, err
	}

	if closeSnapshot {
		_ = snapshot.CloseProcess()
	}
	return snapshot, &module, nil
}

func FindModule(name string, pid uint32) (*Module, error) {
	snapshot, module, err := FindBaseModule(pid, false)
	if err != nil {
		_ = snapshot.CloseProcess()
		return nil, err
	}

	for {
		moduleName := string(module.Name[:])

		if strings.Index(moduleName, name) == 0 {
			_ = snapshot.CloseProcess()
			return module, nil
		}

		err = snapshot.ModuleNext(module)
		if err != nil {
			break
		}
	}

	_ = snapshot.CloseProcess()
	return nil, err
}
