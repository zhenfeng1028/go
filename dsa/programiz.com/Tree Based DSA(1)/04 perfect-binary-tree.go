package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func depth(node *Node) int {
	d := 0
	for node.left != nil {
		d++
		node = node.left
	}
	return d
}

func isPerfectR(node *Node, d int, level int) bool {
	if node == nil {
		return true
	}

	if node.left == nil && node.right == nil {
		return d == level
	}

	if node.left == nil || node.right == nil {
		return false
	}

	return isPerfectR(node.left, d, level+1) && isPerfectR(node.right, d, level+1)
}

func isPerfect(node *Node) bool {
	d := depth(node)
	return isPerfectR(node, d, 0)
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}

	if isPerfect(root) {
		fmt.Println("The tree is a perfect binary tree")
	} else {
		fmt.Println("The tree is not a perfect binary tree")
	}
}
