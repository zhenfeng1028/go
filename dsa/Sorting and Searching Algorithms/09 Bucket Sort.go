package main

import "fmt"

const (
	NBUCKET  = 6  // number of buckets
	INTERVAL = 10 // bucket interval
)

type Bucket []int

func bucketSort(arr []int) {
	// Create buckets
	buckets := make([]Bucket, NBUCKET)

	// Initialize empty buckets
	for i := 0; i < NBUCKET; i++ {
		buckets[i] = make(Bucket, 0)
	}

	// Fill the buckets with with respective elements
	for i := 0; i < len(arr); i++ {
		pos := getBucketIndex(arr[i])
		buckets[pos] = append(buckets[pos], arr[i])
	}

	// Print the buckets along with their elements
	for i := 0; i < NBUCKET; i++ {
		fmt.Printf("Bucket[%d]: ", i)
		printBucket(buckets[i])
		fmt.Println()
	}

	// Sort the elements of each bucket
	for i := 0; i < NBUCKET; i++ {
		insertionSort(buckets[i], len(buckets[i]))
	}

	fmt.Println("-------------")
	fmt.Println("Buckets after sorted")
	for i := 0; i < NBUCKET; i++ {
		fmt.Printf("Bucket[%d]: ", i)
		printBucket(buckets[i])
		fmt.Println()
	}

	cursor := 0
	// Put sorted elements on array
	for i := 0; i < NBUCKET; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[cursor] = buckets[i][j]
			cursor++
		}
	}
}

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

func printBucket(bucket Bucket) {
	for i := 0; i < len(bucket); i++ {
		fmt.Print(bucket[i], " ")
	}
}

func getBucketIndex(value int) int {
	return value / INTERVAL
}

func main() {
	a := []int{42, 33, 37, 52, 32, 47, 51}

	fmt.Println("Initial array:")
	fmt.Println(a)
	fmt.Println("-------------")

	bucketSort(a)
	fmt.Println("-------------")
	fmt.Println("Sorted array:")
	fmt.Println(a)
}
