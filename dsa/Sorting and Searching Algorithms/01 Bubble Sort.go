package main

import "fmt"

func bubbleSort(arr []int, size int) {
	for step := 1; step < size; step++ {
		for i := 0; i < size-step; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
}

func main() {
	var a = []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	bubbleSort(a, len(a))
	fmt.Println(a)
}
