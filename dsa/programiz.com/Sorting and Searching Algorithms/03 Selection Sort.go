package main

import "fmt"

func selectionSort(arr []int, size int) {
	for step := 0; step < size-1; step++ {
		minIdx := step
		for i := step + 1; i < size; i++ {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
		}
		temp := arr[step]
		arr[step] = arr[minIdx]
		arr[minIdx] = temp
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	selectionSort(a, len(a))
	fmt.Println(a)
}
