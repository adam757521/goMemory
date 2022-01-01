package memory

import "syscall"

func callProc(proc *syscall.Proc, params ...uintptr) (uintptr, uintptr, error) {
	r1, r2, err := proc.Call(params...)

	if r1 == 0 {
		return 0, 0, err
	}

	return r1, r2, nil
}
