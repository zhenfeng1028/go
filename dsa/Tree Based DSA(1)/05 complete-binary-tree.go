package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

// Count the number of nodes
func countNumNodes(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + countNumNodes(node.left) + countNumNodes(node.right)
}

// Check if the tree is a complete binary tree
func checkComplete(node *Node, index int, numberNodes int) bool {
	if node == nil {
		return true
	}

	if index >= numberNodes {
		return false
	}

	return checkComplete(node.left, 2*index+1, numberNodes) && checkComplete(node.right, 2*index+2, numberNodes)
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}

	nodeCount := countNumNodes(root)

	if checkComplete(root, 0, nodeCount) {
		fmt.Println("The tree is a complete binary tree")
	} else {
		fmt.Println("The tree is not a complte binary tree")
	}
}
