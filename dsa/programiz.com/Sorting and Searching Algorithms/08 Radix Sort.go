package main

import "fmt"

const digit_num = 10

// Function to get the largest element from an array
func getMax(arr []int, size int) int {
	max := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// Using counting sort to sort the elements in the basis of significant places
func countSort(arr []int, size int, place int) {
	output := make([]int, size)

	count := make([]int, digit_num)

	// Store the count of each element
	for i := 0; i < size; i++ {
		count[(arr[i]/place)%10]++
	}

	// Store the cumulative count of each array
	for i := 1; i < digit_num; i++ {
		count[i] += count[i-1]
	}

	// Place the elements in sorted order
	for i := size - 1; i >= 0; i-- {
		output[count[(arr[i]/place)%10]-1] = arr[i]
		count[(arr[i]/place)%10]--
	}

	// Copy the sorted elements into original array
	for i := 0; i < size; i++ {
		arr[i] = output[i]
	}
}

func radixSort(arr []int, size int) {
	// Get maximum element
	max := getMax(arr, size)

	// Apply counting sort to sort elements based on place value
	for place := 1; max/place > 0; place *= 10 {
		countSort(arr, size, place)
	}
}

func main() {
	a := []int{121, 432, 564, 23, 1, 45, 788}
	radixSort(a, len(a))
	fmt.Println(a)
}
