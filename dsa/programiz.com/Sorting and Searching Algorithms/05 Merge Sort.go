package main

import "fmt"

// Merge two subarrays L and M into arr
func merge(arr []int, p, q, r int) {
	// Create L ← A[p..q] and M ← A[q+1..r]
	n1 := q - p + 1
	n2 := r - q

	L, M := make([]int, n1), make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[p+i]
	}
	for j := 0; j < n2; j++ {
		M[j] = arr[q+1+j]
	}

	// Maintain current index of sub-arrays and main array
	var i, j, k int
	i = 0
	j = 0
	k = p

	// Until we reach end of either L or M, pick smaller elements
	// among L and M and place them in the correct position at A[p..r]
	for i < n1 && j < n2 {
		if L[i] <= M[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = M[j]
			j++
		}
		k++
	}

	// When we run out of elements in either L or M,
	// pick up the remaining elements and put in A[p..r]
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = M[j]
		j++
		k++
	}
}

// Divide the array into two subarrays, sort them and merge them
func mergeSort(arr []int, l, r int) {
	if l < r {
		// m is the point where the array is divided into two subarrays
		m := l + (r-l)/2

		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		// Merge the sorted subarrays
		merge(arr, l, m, r)
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	mergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}
