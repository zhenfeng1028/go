package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"abc", "def", "ghi", "jkl"}

	x := "def"
	i := sort.SearchStrings(a, x)
	fmt.Printf("found %q at index %d in %v\n", x, i, a)

	x = "mno"
	i = sort.SearchStrings(a, x)
	fmt.Printf("%q not found, can be inserted at index %d in %v\n", x, i, a)
}
