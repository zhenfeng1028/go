package main

import (
	"fmt"
)

// Enclosing function: returns a memoized version of f
func Memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)

	return func(n int) int {
		if val, ok := cache[n]; ok {
			fmt.Println("From cache:", n)
			return val
		}

		result := f(n)
		cache[n] = result
		return result
	}
}

func main() {
	// Recursive Fibonacci function (not yet memoized)
	var fib func(int) int

	// Define fib using closure
	fib = Memoize(func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	})

	fmt.Println(fib(10)) // Should compute only once per value
}
