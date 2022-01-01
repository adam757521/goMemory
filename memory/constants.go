package memory

import "syscall"

var (
	user   = syscall.MustLoadDLL("User32.dll")
	kernel = syscall.MustLoadDLL("kernel32.dll")

	kernelReadProcess    = kernel.MustFindProc("ReadProcessMemory")        // BOOL ReadProcessMemory(HANDLE hProcess, LPCVOID lpBaseAddress, LPVOID lpBuffer, SIZE_T nSize, SIZE_T* lpNumberOfBytesRead)
	kernelWriteProcess   = kernel.MustFindProc("WriteProcessMemory")       // BOOL WriteProcessMemory(HANDLE hProcess, LPVOID lpBaseAddress, LPCVOID lpBuffer, SIZE_T nSize, SIZE_T* lpNumberOfBytesWritten)
	kernelOpenProcess    = kernel.MustFindProc("OpenProcess")              // HANDLE OpenProcess(DWORD dwDesiredAccess, BOOL bInheritHandle, DWORD dwProcessId)
	kernelCloseProcess   = kernel.MustFindProc("CloseHandle")              // BOOL CloseHandle(HANDLE hObject)
	kernelHelpCreateTool = kernel.MustFindProc("CreateToolhelp32Snapshot") // HANDLE CreateToolhelp32Snapshot(DWORD dwFlags, DWORD th32ProcessID)
	kernelModuleFirst    = kernel.MustFindProc("Module32First")            // BOOL Module32First(HANDLE hSnapshot, LPMODULEENTRY32 lpme)
	kernelModuleNext     = kernel.MustFindProc("Module32Next")             // BOOL Module32Next(HANDLE hSnapshot, LPMODULEENTRY32 lpme)
	kernelGetProcessId   = kernel.MustFindProc("GetProcessId")             // DWORD GetProcessId(HANDLE hProcess)
	kernelProcessFirst   = kernel.MustFindProc("Process32First")           // BOOL Process32First(HANDLE hSnapshot, LPPROCESSENTRY32 lppe)
	kernelProcessNext    = kernel.MustFindProc("Process32Next")            // BOOL Process32Next(HANDLE hSnapshot, LPPROCESSENTRY32 lppe)

	userGetForegroundWindow = user.MustFindProc("GetForegroundWindow") // HWND GetForegroundWindow()
	userGetKeyState         = user.MustFindProc("GetKeyState")         // SHORT GetKeyState(int nVirtKey)
)

const (
	ProcessAllAccess = 0x1F0FFF
	SnapProcess      = 0x00000002
	SnapModules      = 0x00000010 | 0x00000008
	Nil              = uintptr(0) // Used for Read/Write where the read memory amount is not needed.
)
