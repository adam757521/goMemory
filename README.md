<h1 align="center">goMemory</h1>

<p align="center">
   A go module that supports reading and writing process memory.
   </br>
   Useful in game hacking.
</p>

Examples
--------------

### Read and write process memory ###

```go
package main

import (
	"os"
	"unsafe"

	"github.com/adam757521/goMemory/memory"
)

func main() {
	handle, err := memory.OpenProcess(memory.ProcessAllAccess, false, uint32(os.Getpid()))
	if err != nil {
		println("Could not open current process.")
		return
	}

	myLuckyNumber := 5
	myLuckyNumberAddress := uintptr(unsafe.Pointer(&myLuckyNumber))

	myLuckyNumberReadMemory, err := handle.ReadMemory(myLuckyNumberAddress)
	if err != nil {
		println("Could not read myLuckyNumber address.")
		return
	}

	println("Read", myLuckyNumberReadMemory, "from myLuckyNumber address.")

	err = handle.WriteMemory(myLuckyNumberAddress, 38)
	if err != nil {
		println("Could not write 38 to myLuckyNumber address.")
		return
	}

	println("myLuckyNumber is now", myLuckyNumber) // 38
	
	_ = handle.CloseProcess()
}
```
