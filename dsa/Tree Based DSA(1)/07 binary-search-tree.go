package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

// inorder traversal
func inorder(node *Node) {
	if node != nil {
		// traverse left
		inorder(node.left)
		// traverse self
		fmt.Print(node.data, " -> ")
		// traverse right
		inorder(node.right)
	}
}

// insert a node
func insert(node *Node, key int) *Node {
	// return a new node if the tree is empty
	if node == nil {
		return &Node{data: key}
	}

	// traverse to the right place and insert the node
	if key < node.data {
		node.left = insert(node.left, key)
	} else {
		node.right = insert(node.right, key)
	}

	return node
}

// find the inorder successor
func minValueNode(node *Node) *Node {
	// find the leftmost leaf
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

// delete a node
func delete(node *Node, key int) *Node {
	// return if the tree is empty
	if node == nil {
		return node
	}

	// find the node to be deleted
	if key < node.data {
		node.left = delete(node.left, key)
	} else if key > node.data {
		node.right = delete(node.right, key)
	} else {
		// if the node is with only one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.right
		}

		// if the node has two children
		temp := minValueNode(node.right)

		// place the inorder successor in position of the node to be deleted
		node.data = temp.data

		// delete the inorder successor
		node.right = delete(node.right, temp.data)
	}

	return node
}

func main() {
	var root *Node

	root = insert(root, 8)
	root = insert(root, 3)
	root = insert(root, 1)
	root = insert(root, 6)
	root = insert(root, 7)
	root = insert(root, 10)
	root = insert(root, 14)
	root = insert(root, 4)

	fmt.Print("Inorder traversal: ")
	inorder(root)
	fmt.Printf("\n")

	fmt.Printf("After deleting 3\n")
	delete(root, 3)
	fmt.Print("Inorder traversal: ")
	inorder(root)
}
