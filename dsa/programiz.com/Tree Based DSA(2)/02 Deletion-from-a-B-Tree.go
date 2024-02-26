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

// Return the index of the first key that is greater than or
// equal to key
func (node *BTreeNode) findKey(k int) int {
	idx := 0
	for idx < node.n && node.keys[idx] < k {
		idx++
	}
	return idx
}

// Get predecessor of keys[idx]
func (node *BTreeNode) getPred(idx int) int {
	// Keep moving to the right most node until we reach a leaf
	cur := node.C[idx]
	for !cur.leaf {
		cur = cur.C[cur.n]
	}

	// Return the last key of the leaf
	return cur.keys[cur.n-1]
}

// Get successor of keys[idx]
func (node *BTreeNode) getSucc(idx int) int {
	// Keep moving to the left most node starting from C[idx+1] until we reach a leaf
	cur := node.C[idx+1]
	for !cur.leaf {
		cur = cur.C[0]
	}

	// Return the first key of the leaf
	return cur.keys[0]
}

// Remove the key k from the sub-tree rooted with this node
func (node *BTreeNode) remove(k int) {
	idx := node.findKey(k)

	// The key to be removed is present in this node
	if idx < node.n && node.keys[idx] == k {
		// If the node is a leaf node - removeFromLeaf is called
		// Otherwise, removeFromNonLeaf is called
		if node.leaf {
			node.removeFromLeaf(idx)
		} else {
			node.removeFromNonLeaf(idx)
		}
	} else {
		// If the node is a leaf node, then the key is not present in the tree
		if node.leaf {
			fmt.Printf("The key %d does not exist in the tree\n", k)
			return
		}

		// The key to be removed is present in the subtree rooted with this node
		// The flag indicates whether the key is present in the subtree rooted
		// with the last child of this node
		flag := idx == node.n

		// If the child where the key is supposed to exist has less than t keys,
		// we fill that child
		if node.C[idx].n < node.t {
			node.fill(idx)
		}

		// If the last child has been merged, it must have merged with the previous
		// child and so we recurse on the (idx-1)th child. Else, we recurse on the
		// (idx)th child which now has atleast t keys
		if flag && idx > node.n {
			node.C[idx-1].remove(k)
		} else {
			node.C[idx].remove(k)
		}
	}
}

// Remove the idx-th key from this node - which is a leaf node
func (node *BTreeNode) removeFromLeaf(idx int) {
	// Move all the keys after the idx-th pos one place backward
	for i := idx + 1; i < node.n; i++ {
		node.keys[i-1] = node.keys[i]
	}

	// Reduce the count of keys
	node.n--
}

// Remove the idx-th key from this node - which is a non-leaf node
func (node *BTreeNode) removeFromNonLeaf(idx int) {
	k := node.keys[idx]

	// If the child that precedes k (C[idx]) has atleast t keys,
	// find the predecessor 'pred' of k in the subtree rooted at
	// C[idx]. Replace k by pred. Recursively delete pred in C[idx]
	if node.C[idx].n >= node.t {
		pred := node.getPred(idx)
		node.keys[idx] = pred
		node.C[idx].remove(pred)
	} else if node.C[idx+1].n >= node.t {
		// If the child C[idx] as less than t keys, examine C[idx+1].
		// If C[idx+1] has atleast keys, find the successor 'succ' of k
		// in the subtree rooted at C[idx+1]
		// Replace k by succ
		// Recursively delete succ in C[idx+1]
		succ := node.getSucc(idx)
		node.keys[idx] = succ
		node.C[idx+1].remove(succ)
	} else {
		// If both C[idx] and C[idx+1] has less than t keys,
		// merge k and all of C[idx+1] into C[idx]
		// Now C[idx] contains 2t-1 keys
		// Free C[idx+1] and recursively delete k from C[idx]
		node.merge(idx)
		node.C[idx].remove(k)
	}
}

// fill child C[idx] which has less than t-1 keys
func (node *BTreeNode) fill(idx int) {
	// If the previous child(C[idx-1]) has more than t-1 keys, borrorw
	// a key from that child
	if idx != 0 && node.C[idx-1].n >= node.t {
		node.borrorFromPrev(idx)
	} else if idx != node.n && node.C[idx+1].n >= node.t {
		// If the next child(C[idx+1]) has more than t-1 keys, borrow
		// a key from that child
		node.borrorFromNext(idx)
	} else {
		// Merge C[idx] with its sibling
		// If C[idx] is last child, merge it with its previous sibling
		// Otherwise merge it with its next sibling
		if idx != node.n {
			node.merge(idx)
		} else {
			node.merge(idx - 1)
		}
	}
}

// Borrow a key from C[idx-1] and place it in C[idx]
func (node *BTreeNode) borrorFromPrev(idx int) {
	child := node.C[idx]
	sibling := node.C[idx-1]

	// The last key from C[idx-1] goes up to the parent and key[idx-1]
	// from parent is inserted as the first key in C[idx]. Thus, the
	// sibling loses one key and child gains one key

	// Moving all keys in C[idx] one step ahead
	for i := child.n - 1; i >= 0; i-- {
		child.keys[i+1] = child.keys[i]
	}

	// If C[idx] is not a leaf, move all its child pointers one step ahead
	if !child.leaf {
		for i := child.n; i >= 0; i-- {
			child.C[i+1] = child.C[i]
		}
	}

	// Setting child's first key equal to keys[idx-1] from the current node
	child.keys[0] = node.keys[idx-1]

	// Moving sibling's last child as C[idx]'s first child
	if !child.leaf {
		child.C[0] = sibling.C[sibling.n]
	}

	// Moving the key from the sibling to the parent
	node.keys[idx-1] = sibling.keys[sibling.n-1]

	// Increasing and decreasing the key count of C[idx] and C[idx+1]
	// respectively
	child.n++
	sibling.n--
}

// Borror a key from C[idx+1] and place it in C[idx]
func (node *BTreeNode) borrorFromNext(idx int) {
	child := node.C[idx]
	sibling := node.C[idx+1]

	// keys[idx] is inserted as the last key in C[idx]
	child.keys[child.n] = node.keys[idx]

	// Sibling's first child is inserted as the last child
	// into C[idx]
	if !child.leaf {
		child.C[child.n+1] = sibling.C[0]
	}

	// The first key from sibling is inserted into keys[idx]
	node.keys[idx] = sibling.keys[0]

	// Moving all keys in sibling one step backward
	for i := 1; i < sibling.n; i++ {
		sibling.keys[i-1] = sibling.keys[i]
	}

	// Moving the child pointers one step backward
	if !sibling.leaf {
		for i := 1; i < sibling.n; i++ {
			sibling.C[i-1] = sibling.C[i]
		}
	}

	// Increasing and decreasing the key count of C[idx] and C[idx+1]
	// respectively
	child.n++
	sibling.n--
}

// Merge C[idx] with C[idx+1], C[idx+1] is freed after merging
func (node *BTreeNode) merge(idx int) {
	child := node.C[idx]
	sibling := node.C[idx+1]

	// Pulling a key from the current node and inserting it into (t-1)th
	// position of C[idx]
	child.keys[child.t-1] = node.keys[idx]

	// Coping the keys from C[idx+1] to C[idx] at the end
	for i := 0; i < sibling.n; i++ {
		child.keys[i+child.t] = sibling.keys[i]
	}

	// Coping the child pointers from C[idx+1] to C[idx]
	if !child.leaf {
		for i := 0; i < sibling.n; i++ {
			child.C[i+child.t] = sibling.C[i]
		}
	}

	// Moving all keys after idx in the current node one step backward
	// to fill the gap created by moving keys[idx] to C[idx]
	for i := idx + 1; i < node.n; i++ {
		node.keys[i-1] = node.keys[i]
	}

	// Moving the child pointers after (idx+1) in current node
	// one step backward
	for i := idx + 2; i <= node.n; i++ {
		node.C[i-1] = node.C[i]
	}

	// Updating the key count of child and the current node
	child.n += sibling.n + 1
	node.n--

	// Freeing the memory occupied by sibling
	sibling = nil
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

func (tree *BTree) remove(k int) {
	if tree.root == nil {
		fmt.Println("The tree is empty")
		return
	}

	// Call the remove funtion for root
	tree.root.remove(k)

	// If the root node has 0 keys, make its first child as the new root
	// if it has a child, otherwise set root as NULL
	if tree.root.n == 0 {
		if tree.root.leaf {
			tree.root = nil
		} else {
			tree.root = tree.root.C[0]
		}
	}
}

func main() {
	t := NewBTree(3) // A B-Tree with minimum degree 3

	t.insert(1)
	t.insert(3)
	t.insert(7)
	t.insert(10)
	t.insert(11)
	t.insert(13)
	t.insert(14)
	t.insert(15)
	t.insert(18)
	t.insert(16)
	t.insert(19)
	t.insert(24)
	t.insert(25)
	t.insert(26)
	t.insert(21)
	t.insert(4)
	t.insert(5)
	t.insert(20)
	t.insert(22)
	t.insert(2)
	t.insert(17)
	t.insert(12)
	t.insert(6)

	fmt.Println("Traversal of tree constructed is")
	t.traverse()

	t.remove(6)
	fmt.Printf("\nTraversal of tree after removing 6\n")
	t.traverse()

	t.remove(13)
	fmt.Printf("\nTraversal of tree after removing 13\n")
	t.traverse()

	t.remove(7)
	fmt.Printf("\nTraversal of tree after removing 7\n")
	t.traverse()

	t.remove(4)
	fmt.Printf("\nTraversal of tree after removing 4\n")
	t.traverse()

	t.remove(2)
	fmt.Printf("\nTraversal of tree after removing 2\n")
	t.traverse()

	t.remove(16)
	fmt.Printf("\nTraversal of tree after removing 16\n")
	t.traverse()
}
