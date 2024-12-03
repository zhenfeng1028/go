package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	clone := slices.Clone(numbers)
	fmt.Println(clone)
	clone[2] = 10
	fmt.Println(numbers)
}
