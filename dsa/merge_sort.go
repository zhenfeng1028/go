package main

import (
	"fmt"
)

func mergeSort(nums []int) []int {
	// 分治，两两拆分，一直拆到基础元素才向上递归
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	// 左侧数据递归拆分
	left := mergeSort(nums[0:i])
	// 右侧数据递归拆分
	right := mergeSort(nums[i:])
	// 排序 & 合并
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func main() {
	a := []int{3, 2, 5, 4, 6, 9, 7, 8, 1}
	result := mergeSort(a)
	fmt.Println(result)
}
