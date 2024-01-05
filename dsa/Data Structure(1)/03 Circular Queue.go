package main

import "fmt"

const initSize = 5

type Queue struct {
	items []int
	front int
	rear  int
	size  int
}

func NewQueue() *Queue {
	return &Queue{front: -1, rear: -1}
}

func (q *Queue) IsFull() bool {
	return (q.front == 0 && q.rear == q.size-1) || (q.front == q.rear+1)
}

func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

func (q *Queue) Enqueue(element int) {
	if q.IsFull() {
		fmt.Println("Queue is full")
	} else {
		if q.front == -1 {
			q.front = 0
		}
		q.rear = (q.rear + 1) % q.size
		q.items[q.rear] = element
		fmt.Println("Inserted", element)
	}
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Empty Queue")
		return -1
	} else {
		element := q.items[q.front]
		// Q has only one element,
		// so we reset the queue after deleting it.
		if q.front == q.rear {
			q.front = -1
			q.rear = -1
		} else {
			q.front = (q.front + 1) % q.size
		}
		fmt.Println("Deleted", element)
		return element
	}
}

func (q *Queue) Display() {
	fmt.Println("Front index ->", q.front)
	fmt.Print("Items -> ")
	var i int
	for i = q.front; i != q.rear; i = (i + 1) % q.size {
		fmt.Printf("%d ", q.items[i])
	}
	fmt.Println(q.items[i])
	fmt.Println("Rear index ->", q.rear)
}

func main() {
	q := NewQueue()

	// Dequeue is not possible on empty queue
	q.Dequeue()

	// Enqueue 5 elements
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	// Fails to enqueue because front == 0 && rear == q.size - 1
	q.Enqueue(6)

	q.Display()

	q.Dequeue()

	q.Display()

	q.Enqueue(7)

	q.Display()

	// Fails to enqueue because front == rear + 1
	q.Enqueue(8)
}
