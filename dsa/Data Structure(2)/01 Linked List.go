package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	var head *Node
	one := &Node{value: 1}
	two := &Node{value: 2}
	three := &Node{value: 3}

	one.next = two
	two.next = three
	three.next = nil

	head = one
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}
