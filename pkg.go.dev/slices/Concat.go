package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{4, 5, 6}
	concat := slices.Concat(s1, s2)
	fmt.Println(concat)
}
