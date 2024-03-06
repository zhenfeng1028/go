package main

import "fmt"

const initSize = 5

type Deque struct {
	items []int
	front int
	rear  int
	size  int
}

func NewDeque() *Deque {
	return &Deque{
		front: -1,
		rear:  0,
		size:  initSize,
		items: make([]int, initSize),
	}
}

func (q *Deque) IsFull() bool {
	return (q.front == 0 && q.rear == q.size-1) || (q.front == q.rear+1)
}

func (q *Deque) IsEmpty() bool {
	return q.front == -1
}

func (q *Deque) InsertFront(key int) {
	if q.IsFull() {
		fmt.Println("Overflow")
		return
	}

	if q.front == -1 {
		q.front = 0
		q.rear = 0
	} else if q.front == 0 {
		q.front = q.size - 1
	} else {
		q.front = q.front - 1
	}

	q.items[q.front] = key
}

func (q *Deque) InsertRear(key int) {
	if q.IsFull() {
		fmt.Println("Overflow")
		return
	}

	if q.front == -1 {
		q.front = 0
		q.rear = 0
	} else if q.rear == q.size-1 {
		q.rear = 0
	} else {
		q.rear = q.rear + 1
	}

	q.items[q.rear] = key
}

func (q *Deque) DeleteFront() {
	if q.IsEmpty() {
		fmt.Println("Underflow")
		return
	}

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else if q.front == q.size-1 {
		q.front = 0
	} else {
		q.front = q.front + 1
	}
}

func (q *Deque) DeleteRear() {
	if q.IsEmpty() {
		fmt.Println("Underflow")
		return
	}

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else if q.rear == 0 {
		q.rear = q.size - 1
	} else {
		q.rear = q.rear - 1
	}
}

func (q *Deque) GetFront() int {
	if q.IsEmpty() {
		fmt.Println("Underflow")
		return -1
	}
	return q.items[q.front]
}

func (q *Deque) GetRear() int {
	if q.IsEmpty() {
		fmt.Println("Underflow")
		return -1
	}
	return q.items[q.rear]
}

func main() {
	q := NewDeque()

	fmt.Println("insert element at rear end")
	q.InsertRear(5)
	q.InsertRear(11)

	fmt.Printf("rear element: %d\n", q.GetRear())

	q.DeleteRear()
	fmt.Printf("after deletion of the rear element, the new rear element: %d\n", q.GetRear())

	fmt.Println("insert element at front end")
	q.InsertFront(8)

	fmt.Printf("front element: %d\n", q.GetFront())

	q.DeleteFront()
	fmt.Printf("after deletion of the front element, the new front element: %d\n", q.GetFront())
}
