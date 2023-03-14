package main

import "fmt"

type Callback func(int) int

func getOddNumber(k int, callback Callback) int {
	return 1 + callback(k)
}

// 2x
func double(x int) int {
	return 2 * x
}

// 4x
func quadruple(x int) int {
	return 4 * x
}

func main() {
	fmt.Println(getOddNumber(1, double))
	fmt.Println(getOddNumber(1, quadruple))
	fmt.Println(getOddNumber(1, func(x int) int {
		return 8 * x
	}))
}
