package main

import "fmt"

func swap(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func heapify(s []int, i int) {
	size := len(s)
	// Find the largest among root, left child and right child
	largest := i
	l, r := 2*i+1, 2*i+2
	if l < size && s[l] > s[largest] {
		largest = l
	}
	if r < size && s[r] > s[largest] {
		largest = r
	}
	// Swap and continue heapifying if root is not largest
	if largest != i {
		swap(s, i, largest)
		heapify(s, largest)
	}
}

func insert(s *[]int, newNum int) {
	size := len(*s)
	if size == 0 {
		*s = append(*s, newNum)
	} else {
		*s = append(*s, newNum)
		for i := size/2 - 1; i >= 0; i-- {
			heapify(*s, i)
		}
	}
}

func deleteNode(s *[]int, num int) {
	size := len(*s)
	i := 0
	for i = 0; i < size; i++ {
		if (*s)[i] == num {
			break
		}
	}
	swap(*s, i, size-1)
	*s = (*s)[:size-1]
	for i := size/2 - 1; i >= 0; i-- {
		heapify(*s, i)
	}
}

func printArray(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
}

func main() {
	var heapTree []int

	insert(&heapTree, 3)
	insert(&heapTree, 4)
	insert(&heapTree, 9)
	insert(&heapTree, 5)
	insert(&heapTree, 2)

	fmt.Print("Max-Heap array: ")
	printArray(heapTree)

	deleteNode(&heapTree, 4)

	fmt.Print("After deleting an element: ")
	printArray(heapTree)
}
