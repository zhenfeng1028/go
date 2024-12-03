package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 1, 2, 3}
	fmt.Println(slices.Contains(numbers, 2))
	fmt.Println(slices.Contains(numbers, 4))
}
