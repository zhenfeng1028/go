package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	grow := slices.Grow(numbers, 2)
	fmt.Println(cap(numbers))
	fmt.Println(grow)
	fmt.Println(len(grow))
	fmt.Println(cap(grow))
}
