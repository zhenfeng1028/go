package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 1, 2, 3}
	repeat := slices.Repeat(numbers, 2)
	fmt.Println(repeat)
}
