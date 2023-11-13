package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		// do something
	}()
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Println(string(buf))
}
