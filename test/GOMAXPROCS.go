package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(8)
	now := time.Now()
	for range 10000 {
		var a uint64
		wg.Add(1)
		go func() {
			for range math.MaxInt16 {
				atomic.AddUint64(&a, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("time elapsed:", time.Since(now))
}

// 1 CPU: 1.574640024s
// 4 CPUs: 550.466936ms
// 8 CPUs: 457.705519ms
