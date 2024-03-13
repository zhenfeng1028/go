package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func addToEmpty(last_ref **Node, data int) {
	if *last_ref != nil {
		return
	}

	newNode := &Node{data: data}

	*last_ref = newNode

	// create link to itself
	(*last_ref).next = *last_ref
}

// add node to the front
func addFront(last_ref **Node, data int) {
	// check if the list is empty
	if *last_ref == nil {
		addToEmpty(last_ref, data)
		return
	}

	newNode := &Node{data: data}

	// store the address of the current first node in the newNode
	newNode.next = (*last_ref).next

	// make the newNode as head
	(*last_ref).next = newNode
}

// add node to the end
func addEnd(last_ref **Node, data int) {
	// check if the list is empty
	if *last_ref == nil {
		addToEmpty(last_ref, data)
		return
	}

	newNode := &Node{data: data}

	// store the address of the head node in the newNode
	newNode.next = (*last_ref).next

	// point the current last node to the newNode
	(*last_ref).next = newNode

	// make newNode as the last node
	(*last_ref) = newNode
}

// insert node after a specific node
func addAfter(last_ref **Node, data, item int) {
	// check if the list is empty
	if *last_ref == nil {
		return
	}

	// specific node to find
	var p *Node
	var i int = 0

	for p != (*last_ref).next {
		// assign head to p
		if i == 0 {
			p = (*last_ref).next
			i++
		}
		// if the item is found, place newNode after it
		if p.data == item {
			newNode := &Node{data: data}

			// make the next of current node as the next of newNode
			newNode.next = p.next

			// put newNode to the next of p
			p.next = newNode

			// if p is the last node, make newNode as the last node
			if p == *last_ref {
				*last_ref = newNode
			}
			return
		}
	}

	fmt.Println("The given node is not present in the list")
}

// delete a node
func deleteNode(last_ref **Node, item int) {
	// if list is empty
	if *last_ref == nil {
		return
	}

	// if the list contains only a single node
	if (*last_ref).data == item && (*last_ref).next == *last_ref {
		*last_ref = nil
		return
	}

	temp := *last_ref

	// if last is to be deleted
	if (*last_ref).data == item {
		// find the node before the last node
		for temp.next != *last_ref {
			temp = temp.next
		}

		// point temp node to the next of last i.e. first node
		temp.next = (*last_ref).next

		// make temp as the last node
		*last_ref = temp
		return
	}

	// travel to the node to be deleted
	for temp.next != *last_ref && temp.next.data != item {
		temp = temp.next
	}

	// if node to be deleted is found
	if temp.next.data == item {
		temp.next = temp.next.next
	}
}

func traverse(last *Node) {
	if last == nil {
		fmt.Println("The list is empty")
		return
	}

	var p *Node
	var i int
	for p != last.next {
		if i == 0 {
			p = last.next
			i++
		}
		fmt.Print(p.data, " ")
		p = p.next
	}
	fmt.Println()
}

func main() {
	var last *Node

	addToEmpty(&last, 6)
	addEnd(&last, 8)
	addFront(&last, 2)
	addAfter(&last, 10, 2)

	traverse(last)

	deleteNode(&last, 8)

	traverse(last)
}
