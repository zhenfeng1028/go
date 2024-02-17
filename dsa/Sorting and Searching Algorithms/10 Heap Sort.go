package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func heapify(arr []int, n int, i int) {
	// Find largest among root, left child and right child
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// Swap and continue heapifying if root is not largest
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int, n int) {
	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Heap sort
	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)

		// Heapify root element to get highest element at root again
		heapify(arr, i, 0)
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	heapSort(a, len(a))
	fmt.Println(a)
}
