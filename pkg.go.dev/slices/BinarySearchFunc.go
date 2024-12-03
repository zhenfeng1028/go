package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 55},
		{"Bob", 24},
		{"Gopher", 13},
	}
	n, found := slices.BinarySearchFunc(people, Person{"Bob", 0}, func(a, b Person) int {
		return strings.Compare(a.Name, b.Name)
	})
	fmt.Println("Bob:", n, found)
}
