package main

import "fmt"

const initSize = 10

type Stack struct {
	items []int
	top   int
	num   int
	size  int
}

func NewStack() *Stack {
	s := &Stack{}
	s.top = -1
	s.size = initSize
	s.items = make([]int, initSize)
	return s
}

func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(item int) {
	if s.IsFull() {
		fmt.Println("STACK FULL")
	} else {
		s.top++
		s.items[s.top] = item
		s.num++
	}
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("STACK EMPTY")
	} else {
		fmt.Println("item popped =", s.items[s.top])
		s.top--
		s.num--
	}
}

func (s *Stack) Print() {
	fmt.Print("Stack: ")
	for i := 0; i < s.num; i++ {
		fmt.Printf("%d ", s.items[i])
	}
	fmt.Println()
}

func main() {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	s.Print()

	s.Pop()

	fmt.Println("After popping out")
	s.Print()
}
