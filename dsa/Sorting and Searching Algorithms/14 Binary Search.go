package main

import "fmt"

func binarySearch(arr []int, x int, low, high int) int {
	// Repeat until the pointers low and high meet each other
	if low <= high {
		mid := low + (high-low)/2

		// If found at mid, then return it
		if arr[mid] == x {
			return mid
		}

		// Search the left half
		if arr[mid] > x {
			return binarySearch(arr, x, low, mid-1)
		}

		// Search the right half
		return binarySearch(arr, x, mid+1, high)
	}
	return -1
}

func main() {
	a := []int{3, 4, 5, 6, 7, 8, 9}
	x := 4
	idx := binarySearch(a, x, 0, len(a)-1)
	if idx == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Println("Element is found at index", idx)
	}
}
