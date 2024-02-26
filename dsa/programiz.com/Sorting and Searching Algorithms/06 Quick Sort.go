package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Rearrange array (find the partition point)
func partition(arr []int, low, high int) int {
	// select the rightmost element as pivot
	pivot := arr[high]

	// pointer for greater element
	i := (low - 1)

	// traverse each element of the array
	// compare them with the pivot
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			// if element smaller than pivot is found
			// swap it with the greater element pointed by i
			i++

			// swap element at i with element at j
			swap(arr, i, j)
		}
	}

	// swap pivot with the greater element at i
	swap(arr, i+1, high)

	// return the partition point
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// find the pivot element such that
		// elements smaller than pivot are on left of pivot
		// elements greater than pivot are on righ of pivot
		pi := partition(arr, low, high)

		// recursive call on the left of pivot
		quickSort(arr, low, pi-1)

		// recursive call on the right of pivot
		quickSort(arr, pi+1, high)
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
