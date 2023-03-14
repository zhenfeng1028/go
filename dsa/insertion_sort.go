package main

import "fmt"

func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			}
		}
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	insertionSort(a)
	fmt.Println(a)
}
