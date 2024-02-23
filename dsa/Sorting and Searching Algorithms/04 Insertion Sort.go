package main

import "fmt"

func insertionSort(arr []int, size int) {
	for step := 1; step < size; step++ {
		key := arr[step]
		i := step - 1

		// Compare key with each element on the left of it
		// until an element smaller than it is found.
		// For descending order, change key < arr[i] to key > arr[i].
		for i >= 0 && key < arr[i] {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	insertionSort(a, len(a))
	fmt.Println(a)
}
