package main

import "fmt"

// 如果pivot选的是左边的元素，那么就从右指针指向的元素开始和pivot比较
// 当出现比pivot小的元素的时候，就将该元素赋值给左指针指向的元素，并产生交替
// 左指针指向的元素开始和pivot比较，当出现比pivot大的元素的时候，就将该元素赋值给右指针指向的元素，并产生交替
// 如此往复，直至左指针和右指针指向同一个元素，至此pivot对应的元素位置已确定
// 分别对左右子序列重复以上步骤

func quickSort(arr []int, L, R int) {
	if L >= R {
		return
	}
	left, right := L, R
	pivot := arr[left]
	for left < right {
		for left < right && pivot <= arr[right] {
			right--
		}
		if left < right {
			arr[left] = arr[right]
		}
		for left < right && pivot >= arr[left] {
			left++
		}
		if left < right {
			arr[right] = arr[left]
		}
		if left == right {
			arr[left] = pivot
		}
	}

	quickSort(arr, L, left-1)
	quickSort(arr, right+1, R)
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	quickSort(a, 0, 8)
	fmt.Println(a)
}
