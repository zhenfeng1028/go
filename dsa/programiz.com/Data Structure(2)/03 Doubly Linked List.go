package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

// insert a node at the front
func insertFront(head_ref **Node, data int) {
	newNode := &Node{data: data}
	newNode.next = *head_ref
	newNode.prev = nil

	// previous of head (now head is the second node) is newNode
	if *head_ref != nil {
		(*head_ref).prev = newNode
	}

	// head points to newNode
	*head_ref = newNode
}

// insert a node after a specific node
func insertAfter(prevNode *Node, data int) {
	// check if previous node is null
	if prevNode == nil {
		fmt.Println("previous node cannot be null")
		return
	}

	newNode := &Node{data: data}

	newNode.next = prevNode.next
	newNode.prev = prevNode

	// set next of prev node to newNode
	prevNode.next = newNode

	// set prev of newNode's next to newNode
	if newNode.next != nil {
		newNode.next.prev = newNode
	}
}

// insert a node at the end of the list
func insertEnd(head_ref **Node, data int) {
	newNode := &Node{data: data}

	newNode.next = nil

	// store the head node temporarily (for later use)
	temp := *head_ref

	// if the linked list is empty, make the newNode as head node
	if *head_ref == nil {
		newNode.prev = nil
		*head_ref = newNode
		return
	}

	// if the linked list is not empty, traverse to the end of the linked list
	for temp.next != nil {
		temp = temp.next
	}

	// now, the last node of the linked list is temp

	// assign next of the last node (temp) to newNode
	temp.next = newNode

	// assign prev of newNode to temp
	newNode.prev = temp
}

// delete a node from the doubly linked list
func deleteNode(head_ref **Node, del_node *Node) {
	// if head or del_node is null, deletion is not possible
	if *head_ref == nil || del_node == nil {
		return
	}

	// if del_node is head node, point the head pointer to the next of del_node
	if *head_ref == del_node {
		*head_ref = del_node.next
	}

	// if del_node is not at the last node, point the prev of node next to del_node to the previous of del_node
	if del_node.next != nil {
		del_node.next.prev = del_node.prev
	}

	// if del_node is not the first node, point the next of the previous node to the next node of del_node
	if del_node.prev != nil {
		del_node.prev.next = del_node.next
	}
}

// print the doubly linked list
func displayList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " -> ")
		head = head.next
	}
	fmt.Println()
}

func main() {
	var head *Node

	insertEnd(&head, 5)
	insertFront(&head, 1)
	insertFront(&head, 6)
	insertEnd(&head, 9)

	// insert 11 after head
	insertAfter(head, 11)

	// insert 15 after the seond node
	insertAfter(head.next, 15)

	displayList(head)

	// delete the last node
	deleteNode(&head, head.next.next.next.next.next)

	displayList(head)
}
