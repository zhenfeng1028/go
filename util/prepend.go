package main

import "fmt"

func main() {
	arr := []int{}
	arr = append([]int{1}, arr...)
	arr = append([]int{2}, arr...)
	arr = append([]int{3}, arr...)
	fmt.Println(arr)
}
