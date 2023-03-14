package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}

// func main() {
// 	hello()
// 	fmt.Println("main goroutine done!")
// }

func main() {
	go hello()
	fmt.Println("main goroutine done!")
}

/*
func main() {
	go hello()
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}
*/
