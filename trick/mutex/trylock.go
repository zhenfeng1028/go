package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const mutexLocked = 1 << iota

// The mutex which supports try-locking.
type mutex struct {
	sync.Mutex
}

// TryLock acquires the lock only if it is free at the time of invocation.
func (tl *mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&tl.Mutex)), 0, mutexLocked)
}

func SimpleTest() {
	var m mutex
	m.Lock()
	time.Sleep(1 * time.Second)
	if !m.TryLock() {
		fmt.Println("not get lock")
	}
	m.Unlock()
	if m.TryLock() {
		fmt.Println("get lock")
	}
	m.Unlock()
}

func ConcurrentTest() {
	m := &mutex{}
	cnt := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(m *mutex, wg *sync.WaitGroup, cntPtr *int32) {
			for {
				if m.TryLock() {
					*cntPtr = *cntPtr + 1
					m.Unlock()
					wg.Done()
					break
				} else {
					runtime.Gosched()
				}
			}
		}(m, wg, &cnt)
	}
	wg.Wait()
	if cnt != 1000 {
		fmt.Println("count error concurrency")
	}
}

func main() {
	SimpleTest()
	ConcurrentTest()
}
