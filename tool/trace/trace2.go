package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}

// go build trace2.go
// GODEBUG=schedtrace=1000 ./trace2
