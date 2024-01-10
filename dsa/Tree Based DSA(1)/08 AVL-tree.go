package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
	height      int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// caculate height
func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

// left rotate
func leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left
	y.left = x
	x.right = T2
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1
	return y
}

// right rotate
func rightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right
	x.right = y
	y.left = T2
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x
}

// get the balance factor of each node
func getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

// insert a node
func insertNode(node *Node, data int) *Node {
	// find the correct position and insert the node
	if node == nil {
		return &Node{data: data, height: 1}
	}
	if data < node.data {
		node.left = insertNode(node.left, data)
	} else if data > node.data {
		node.right = insertNode(node.right, data)
	} else {
		return node // no insert?
	}

	// update the balance factor of each node and balance the tree
	node.height = max(height(node.left), height(node.right)) + 1
	balanceFactor := getBalanceFactor(node)
	if balanceFactor > 1 {
		if data < node.left.data {
			return rightRotate(node)
		} else if data > node.left.data {
			node.left = leftRotate(node.left)
			return rightRotate(node)
		}
	}
	if balanceFactor < -1 {
		if data > node.right.data {
			return leftRotate(node)
		} else if data < node.right.data {
			node.right = rightRotate(node.right)
			return leftRotate(node)
		}
	}

	return node
}

// inorder successor
func nodeWithMinValue(node *Node) *Node {
	if node == nil {
		return node
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

// delete a node
func deleteNode(node *Node, data int) *Node {
	// find the node and delete it
	if node == nil {
		return node
	}
	if data < node.data {
		node.left = deleteNode(node.left, data)
	} else if data > node.data {
		node.right = deleteNode(node.right, data)
	} else {
		if node.left == nil || node.right == nil {
			temp := node.left
			if node.right != nil {
				temp = node.right
			}
			if temp == nil {
				node = nil
			} else {
				node = temp
			}
		} else {
			temp := nodeWithMinValue(node.right)
			node.data = temp.data
			node.right = deleteNode(node.right, temp.data)
		}
	}

	if node == nil {
		return node
	}

	// update the balance factor of each node and balance the tree
	node.height = max(height(node.left), height(node.right)) + 1
	balanceFactor := getBalanceFactor(node)
	if balanceFactor > 1 {
		if getBalanceFactor(node.left) >= 0 {
			return rightRotate(node)
		} else {
			node.left = leftRotate(node.left)
			return rightRotate(node)
		}
	}
	if balanceFactor < -1 {
		if getBalanceFactor(node.right) <= 0 {
			return leftRotate(node)
		} else {
			node.right = rightRotate(node.right)
			return leftRotate(node)
		}
	}

	return node
}

func printTree(node *Node, indent string, last bool) {
	if node != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "    "
		} else {
			fmt.Print("L----")
			indent += "|   "
		}
		fmt.Println(node.data)
		printTree(node.left, indent, false)
		printTree(node.right, indent, true)
	}
}

func main() {
	var root *Node
	root = insertNode(root, 33)
	root = insertNode(root, 13)
	root = insertNode(root, 53)
	root = insertNode(root, 9)
	root = insertNode(root, 21)
	root = insertNode(root, 61)
	root = insertNode(root, 8)
	root = insertNode(root, 11)
	printTree(root, "", true)
	root = deleteNode(root, 13)
	fmt.Println("After deleting")
	printTree(root, "", true)
}
