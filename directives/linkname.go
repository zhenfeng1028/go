package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

//go:linkname readgstatus runtime.readgstatus
//go:nosplit
func readgstatus(gp unsafe.Pointer) uint32

//go:linkname allgs runtime.allgs
var allgs []unsafe.Pointer

type mutex struct {
	key uintptr
}

//go:linkname allglock runtime.allglock
var allglock mutex

//go:linkname lock runtime.lock
func lock(l *mutex)

//go:linkname unlock runtime.unlock
func unlock(l *mutex)

const (
	_Grunnable = 1
	_Grunning  = 2
)

func NumRunnableGoroutine() (num int) {
	lock(&allglock)
	for _, g := range allgs {
		if readgstatus(g)&^0x1000 == _Grunnable {
			num++
		}
	}
	unlock(&allglock)

	return
}

func NumRunningGoroutine() (num int) {
	lock(&allglock)
	for _, g := range allgs {
		if readgstatus(g)&^0x1000 == _Grunning {
			num++
		}
	}
	unlock(&allglock)

	return
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 模拟耗时运算
			}
		}()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("All:", runtime.NumGoroutine(), "runnable:", NumRunnableGoroutine(), "running", NumRunningGoroutine())
}
