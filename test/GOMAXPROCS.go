package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	a, b, c int64
	wg      sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(2)
	now := time.Now()
	wg.Add(3)
	go func() {
		for range 1000000 {
			a++
		}
		wg.Done()
	}()

	go func() {
		for range 1000000 {
			b++
		}
		wg.Done()
	}()

	go func() {
		for range 1000000 {
			c++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(a + b + c)
	fmt.Println(time.Since(now))
}

// 1 CPU: 4.9691ms
// 2 CPUs: 3.7993ms
// 3 CPUs: 2.1373ms
