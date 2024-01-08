package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

// Preorder traversal
func preorderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.data, " -> ")
	preorderTraversal(node.left)
	preorderTraversal(node.right)
}

// Inorder traversal
func inorderTraversal(node *Node) {
	if node == nil {
		return
	}

	inorderTraversal(node.left)
	fmt.Print(node.data, " -> ")
	inorderTraversal(node.right)
}

// Postorder traversal
func postorderTraversal(node *Node) {
	if node == nil {
		return
	}

	postorderTraversal(node.left)
	postorderTraversal(node.right)
	fmt.Print(node.data, " -> ")
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 12}
	root.right = &Node{data: 9}
	root.left.left = &Node{data: 5}
	root.left.right = &Node{data: 6}

	fmt.Printf("Preorder traversal: ")
	preorderTraversal(root)

	fmt.Printf("\nInorder traversal: ")
	inorderTraversal(root)

	fmt.Printf("\nPostorder traversal: ")
	postorderTraversal(root)
}
