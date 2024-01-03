package main

import "fmt"

const SIZE = 5

type Queue struct {
	items [SIZE]int
	front int
	rear  int
}

func NewQueue() *Queue {
	return &Queue{front: -1, rear: -1}
}

func (q *Queue) IsFull() bool {
	if q.front == 0 && q.rear == SIZE-1 {
		return true
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	if q.front == -1 {
		return true
	}
	return false
}

func (q *Queue) Enqueue(element int) {
	if q.IsFull() {
		fmt.Println("Queue is full")
	} else {
		if q.front == -1 {
			q.front = 0
		}
		q.rear++
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
		if q.front >= q.rear {
			q.front = -1
			q.rear = -1 /* Q has only one element, so we reset the queue after deleting it. */
		} else {
			q.front++
		}
		fmt.Println("Deleted", element)
		return element
	}
}

func (q *Queue) Display() {
	fmt.Println("Front index ->", q.front)
	fmt.Print("Items -> ")
	for i := q.front; i <= q.rear; i++ {
		fmt.Printf("%d ", q.items[i])
	}
	fmt.Println()
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

	// 6th element can't be added to because the queue is full
	q.Enqueue(6)

	q.Display()

	// Dequeue removes element entered first i.e. 1
	q.Dequeue()

	// Now we have just 4 elements
	q.Display()
}
