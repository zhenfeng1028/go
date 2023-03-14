package main

import (
	"fmt"
	"math"
)

func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func selection(arr []int, left int) {
	minValue := math.MaxInt
	minIndex := 0
	for i := left; i < len(arr); i++ {
		if arr[i] < minValue {
			minIndex = i
			minValue = arr[i]
		}
	}
	swap(arr, left, minIndex)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		selection(arr, i)
	}
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	selectionSort(a)
	fmt.Println(a)
}
