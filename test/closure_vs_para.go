package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}
	time.Sleep(2 * time.Second)
}

// https://stackoverflow.com/questions/30183669/passing-parameters-to-function-closure

// The difference between using a closure vs using a function parameter has to do with
// sharing the same variable vs getting a copy of the value.
