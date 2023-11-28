package main

import (
	"fmt"
	"runtime"
	"time"
)

func getGoroutineID() uint64 {
	var buf [64]byte
	runtime.Stack(buf[:], false)
	var id uint64
	fmt.Sscanf(string(buf[:]), "goroutine %d", &id)
	return id
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("goroutine", getGoroutineID())
		}()
	}
	time.Sleep(100 * time.Millisecond)
}
