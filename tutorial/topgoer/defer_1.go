package main

import "fmt"

func main() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}
