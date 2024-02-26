package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func isFullBinaryTree(node *Node) bool {
	// checking for emptiness
	if node == nil {
		return true
	}

	// checking for the presence of children
	if node.left == nil && node.right == nil {
		return true
	}

	if node.left != nil && node.right != nil {
		return isFullBinaryTree(node.left) && isFullBinaryTree(node.right)
	}

	return false
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.left.right.left = &Node{data: 6}
	root.left.right.right = &Node{data: 7}

	if isFullBinaryTree(root) {
		fmt.Println("The tree is a full binary tree")
	} else {
		fmt.Println("The tree is not a full binary tree")
	}
}
