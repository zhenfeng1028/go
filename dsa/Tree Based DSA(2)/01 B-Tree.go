package main

import "fmt"

type BTreeNode struct {
	keys []int        // An array of keys
	t    int          // Minimun degree
	C    []*BTreeNode // An array of child pointers
	n    int          // Current number of keys
	leaf bool         // Is true when node is leaf. Otherwise false
}

type BTree struct {
	root *BTreeNode // Pointer to root node
	t    int        // Minimum degree
}

func NewBTreeNode(_t int, _leaf bool) *BTreeNode {
	node := &BTreeNode{}

	// Copy the given minimum degree and leaf property
	node.t = _t
	node.leaf = _leaf

	// Allocate memorty for maximum number of possible keys
	// and child pointers
	node.keys = make([]int, 2*node.t-1)
	node.C = make([]*BTreeNode, 2*node.t)

	// Initialize the number of keys as 0
	node.n = 0

	return node
}

// Traverse all nodes in subtree rooted with this node
func (node *BTreeNode) traverse() {
	// There are n keys and n+1 chilren, traverse through n
	// keys and first n children
	var i int
	for i = 0; i < node.n; i++ {
		// If this is not leaf, then before printing keys[i],
		// traverse the subtree rooted with child C[i].
		if !node.leaf {
			node.C[i].traverse()
		}
		fmt.Print(" ", node.keys[i])
	}
	// Print the subtree rooted with the last child
	if !node.leaf {
		node.C[i].traverse()
	}
}

// Search key k in subtree rooted with this node
func (node *BTreeNode) search(k int) *BTreeNode {
	// Find the first key greater than of equal to k
	i := 0
	for i < node.n && k > node.keys[i] {
		i++
	}

	// If the found key is equal to k, return this node
	if node.keys[i] == k {
		return node
	}

	// If the key is not found here and this is a leaf node
	if node.leaf {
		return nil
	}

	// Go to the appropriate child
	return node.C[i].search(k)
}

// Insert a new key in this node. The assumption is,
// the node must be non-full when this function is called
func (node *BTreeNode) insertNonFull(k int) {
	// Initialize index as index of rightmost element
	i := node.n - 1

	// If this is a leaf node
	if node.leaf {
		// The following loop does two things
		// a) Find the location of new key to be inserted
		// b) Move all greater keys one space ahead
		for i >= 0 && node.keys[i] > k {
			node.keys[i+1] = node.keys[i]
			i--
		}

		// Insert the new key at found location
		node.keys[i+1] = k
		node.n = node.n + 1
	} else { // If this node if not leaf
		// Find the child which is going to have the new key
		for i >= 0 && node.keys[i] > k {
			i--
		}

		// See if the found child is full
		if node.C[i+1].n == 2*node.t-1 {
			// If the child is full, then split it
			node.splitChild(i+1, node.C[i+1])

			// After split, the middle key of C[i] goes up and
			// C[i] is splitted into two. See which of the two
			// is going to have the new key
			if node.keys[i+1] < k {
				i++
			}
		}
		node.C[i+1].insertNonFull(k)
	}
}

// Split the child y of this node. i is index of y in child
// array C[]. The child y must be full when this function is called
func (node *BTreeNode) splitChild(i int, y *BTreeNode) {
	// Create a new node which is going to store (t-1) keys
	// of y. Note that y has currently (2t-1) keys.
	z := NewBTreeNode(y.t, y.leaf)
	z.n = node.t - 1

	// Copy the last (t-1) keys of y to z
	for j := 0; j < node.t-1; j++ {
		z.keys[j] = y.keys[node.t+j]
	}

	// Copy the last t children of y to z
	if !y.leaf {
		for j := 0; j < node.t; j++ {
			z.C[j] = y.C[node.t+j]
		}
	}

	// Reduce the number of keys in y
	y.n = node.t - 1

	// Since this node is going to have a new node,
	// move all right children one space ahead
	for j := node.n; j >= i+1; j-- {
		node.C[j+1] = node.C[j]
	}

	// Link the new child to this node
	node.C[i+1] = z

	// A key of y will move to this node. Find the location
	// of new key and move all greater keys one space ahead
	for j := node.n - 1; j >= i; j-- {
		node.keys[j+1] = node.keys[j]
	}

	// Copy the middle key of y to this node
	node.keys[i] = y.keys[node.t-1]

	// Increment count of keys in this node
	node.n = node.n + 1
}

// Initialize tree as empty
func NewBTree(_t int) *BTree {
	return &BTree{t: _t}
}

// Traverse the tree
func (tree *BTree) traverse() {
	if tree.root != nil {
		tree.root.traverse()
	}
}

// Search a key in this tree
func (tree *BTree) search(k int) *BTreeNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.search(k)
}

// Insert a new key in the B-Tree
func (tree *BTree) insert(k int) {
	// If tree is empty
	if tree.root == nil {
		// Allocate memory for root
		tree.root = NewBTreeNode(tree.t, true)
		tree.root.keys[0] = k // Insert key
		tree.root.n = 1       // Update number of keys in root
	} else { // If tree is not empty
		// If root is full, then tree grows in height
		if tree.root.n == 2*tree.t-1 {
			// Allocate memory for new root
			s := NewBTreeNode(tree.t, false)

			// Make old root as child of new root
			s.C[0] = tree.root

			// Split old root and move 1 key to the new root
			s.splitChild(0, tree.root)

			// New root has two children now. Decide which of
			// the two children is going to have new key
			i := 0
			if s.keys[0] < k {
				i++
			}
			s.C[i].insertNonFull(k)

			// Change root
			tree.root = s
		} else { // If root is not full, call insertNonFull for root
			tree.root.insertNonFull(k)
		}
	}
}

func main() {
	t := NewBTree(2) // A B-Tree with minimum degree 2
	t.insert(8)
	t.insert(9)
	t.insert(10)
	t.insert(11)
	t.insert(15)
	t.insert(20)
	t.insert(17)

	fmt.Print("Traversal of the constructed tree is")
	t.traverse()

	k := 6
	if t.search(k) != nil {
		fmt.Printf("\n%d is present", k)
	} else {
		fmt.Printf("\n%d is not present", k)
	}

	k = 15
	if t.search(k) != nil {
		fmt.Printf("\n%d is present", k)
	} else {
		fmt.Printf("\n%d is not present", k)
	}
}
