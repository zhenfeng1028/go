package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func traversalPreOrder(node *Node) {
	if node != nil {
		fmt.Print(node.data, " -> ")
		traversalPreOrder(node.left)
		traversalPreOrder(node.right)
	}
}

func traversalInOrder(node *Node) {
	if node != nil {
		traversalInOrder(node.left)
		fmt.Print(node.data, " -> ")
		traversalInOrder(node.right)
	}
}

func traversalPostOrder(node *Node) {
	if node != nil {
		traversalPostOrder(node.left)
		traversalPostOrder(node.right)
		fmt.Print(node.data, " -> ")
	}
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 12}
	root.right = &Node{data: 9}
	root.left.left = &Node{data: 5}
	root.left.right = &Node{data: 6}

	fmt.Printf("Preorder traversal: ")
	traversalPreOrder(root)

	fmt.Printf("\nInorder traversal: ")
	traversalInOrder(root)

	fmt.Printf("\nPostorder traversal: ")
	traversalPostOrder(root)
}
