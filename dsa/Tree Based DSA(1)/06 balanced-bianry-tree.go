package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// check height balance
func checkHeightBalance(node *Node, height *int) bool {
	if node == nil {
		*height = 0
		return true
	}

	leftHeight, rightHeight := 0, 0

	l := checkHeightBalance(node.left, &leftHeight)
	r := checkHeightBalance(node.right, &rightHeight)

	if leftHeight > rightHeight {
		*height = leftHeight + 1
	} else {
		*height = rightHeight + 1
	}

	if abs(leftHeight, rightHeight) >= 2 {
		return false
	} else {
		return l && r
	}
}

func main() {
	var height int

	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}

	if checkHeightBalance(root, &height) {
		fmt.Println("The tree is balanced, height =", height)
	} else {
		fmt.Println("The tree is not balanced")
	}
}
