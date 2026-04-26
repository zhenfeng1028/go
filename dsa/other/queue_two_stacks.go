package main

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) move() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			n := len(q.in)
			q.out = append(q.out, q.in[n-1])
			q.in = q.in[:n-1]
		}
	}
}

func (q *MyQueue) Pop() int {
	q.move()
	if len(q.out) == 0 {
		panic("pop from empty queue")
	}
	n := len(q.out)
	x := q.out[n-1]
	q.out = q.out[:n-1]
	return x
}

func (q *MyQueue) Peek() int {
	q.move()
	if len(q.out) == 0 {
		panic("peek from empty queue")
	}
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func main() {
	q := &MyQueue{}
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())  // 1
	fmt.Println(q.Pop())   // 1
	fmt.Println(q.Empty()) // false
	fmt.Println(q.Pop())   // 2
	fmt.Println(q.Empty()) // true
}
