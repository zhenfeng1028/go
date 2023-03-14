package main

import "fmt"

func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func bubble(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			swap(arr, i-1, i)
		}
	}
}

func bubbleSort(arr []int) {
	for i := len(arr); i > 1; i-- {
		bubble(arr)
	}
}

func main() {
	var a = []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	bubbleSort(a)
	fmt.Println(a)
}
