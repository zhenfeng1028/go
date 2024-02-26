package main

import "fmt"

func shellSort(arr []int, n int) {
	// Rearrange elements at each n/2, n/4, n/8, ... intervals
	for interval := n / 2; interval > 0; interval /= 2 {
		for i := interval; i < n; i += 1 {
			temp := arr[i]
			j := 0
			for j = i; j >= interval && arr[j-interval] > temp; j -= interval {
				arr[j] = arr[j-interval]
			}
			arr[j] = temp
		}
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	shellSort(a, len(a))
	fmt.Println(a)
}
