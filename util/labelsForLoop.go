package main

import (
	"fmt"
)

func main() {
outside: // declare the label
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i < j {
				break outside // break to that label
			}
			fmt.Println(i, j)
		}
	}

	// prints 0 0

}
