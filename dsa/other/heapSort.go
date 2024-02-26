package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() {
	nums := []int{3, 7, 5, 4, 2, 1, 8, 6, 9}
	h := &hp{sort.IntSlice{}}
	for _, num := range nums {
		heap.Push(h, num)
		fmt.Println(h.IntSlice)
	}
	for i := 0; i < len(nums); i++ {
		x := heap.Pop(h).(int)
		fmt.Printf("%d ", x)
	}
}
