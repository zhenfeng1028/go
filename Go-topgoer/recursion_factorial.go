package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial((n - 1))
}

func main() {
	var i int = 7
	fmt.Println("Factorial of %d is %d\n", i, factorial(i))
}
