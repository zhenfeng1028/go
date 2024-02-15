package main

import "fmt"

func countSort(arr []int, size int) {
	output := make([]int, size)
	max := arr[0]

	// Find the largest element of the array
	for i := 1; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// Initialize count array with length max+1
	count := make([]int, max+1)

	// Store the count of each element
	for i := 0; i < size; i++ {
		count[arr[i]]++
	}

	// Store the cummulative count of each array
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// Find the index of each element of the original array in count array, and
	// place the elements in output array
	for i := size - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	// Copy the sorted elements into original array
	for i := 0; i < size; i++ {
		arr[i] = output[i]
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	countSort(a, len(a))
	fmt.Println(a)
}
