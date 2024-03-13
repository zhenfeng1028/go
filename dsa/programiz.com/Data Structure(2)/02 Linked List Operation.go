package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func insertAtBeginning(head_ref **Node, data int) {
	newNode := &Node{data: data}

	newNode.next = *head_ref

	*head_ref = newNode // move head to new node
}

func insertAfter(prevNode *Node, data int) {
	if prevNode == nil {
		fmt.Println("the given previous node cannot be null")
		return
	}

	newNode := &Node{data: data}

	newNode.next = prevNode.next

	prevNode.next = newNode
}

func insertAtEnd(head_ref **Node, data int) {
	newNode := &Node{data: data}

	if *head_ref == nil {
		*head_ref = newNode
		return
	}

	last := *head_ref
	for last.next != nil {
		last = last.next
	}

	last.next = newNode
}

func deleteNode(head_ref **Node, key int) {
	temp := *head_ref
	if temp != nil && temp.data == key {
		*head_ref = temp.next
		return
	}

	// find the key to be deleted
	prev := &Node{}
	for temp != nil && temp.data != key {
		prev = temp
		temp = temp.next
	}

	// if the key is not present
	if temp == nil {
		return
	}

	// remove the node
	prev.next = temp.next
}

func searchNode(head *Node, key int) bool {
	current := head
	for current != nil {
		if current.data == key {
			return true
		}
		current = current.next
	}
	return false
}

func sortLinkedList(head *Node) {
	current := head
	var index *Node
	var temp int
	if head == nil {
		return
	} else {
		for current != nil {
			// index points to the node next to current
			index = current.next
			for index != nil {
				if current.data > index.data {
					temp = current.data
					current.data = index.data
					index.data = temp
				}
				index = index.next
			}
			current = current.next
		}
	}
}

func printList(node *Node) {
	for node != nil {
		fmt.Print(node.data, " ")
		node = node.next
	}
	fmt.Println()
}

func main() {
	var head *Node

	insertAtEnd(&head, 1)
	insertAtBeginning(&head, 2)
	insertAtBeginning(&head, 3)
	insertAtEnd(&head, 4)
	insertAfter(head.next, 5)

	fmt.Print("Linked List: ")
	printList(head)

	deleteNode(&head, 5)
	fmt.Print("After deleting an element: ")
	printList(head)

	item_to_find := 5
	if searchNode(head, item_to_find) {
		fmt.Println(item_to_find, "is found")
	} else {
		fmt.Println(item_to_find, "is not found")
	}

	sortLinkedList(head)
	fmt.Print("Sorted List: ")
	printList(head)
}
