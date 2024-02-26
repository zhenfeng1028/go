package main

import "fmt"

func search(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{2, 4, 0, 1, 9}
	x := 1
	idx := search(a, x)
	if idx == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Println("Element is found at index", idx)
	}
}
