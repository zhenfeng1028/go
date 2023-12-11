// 找出数组中和为给定值的两个元素的下标
// 例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
package main

import "fmt"

func main() {
	a := [5]int{1, 3, 5, 8, 7}
	MyTest(a, 8)
}

func MyTest(arr [5]int, target int) {
	for i := 0; i < len(arr); i++ {
		other := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == other {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
