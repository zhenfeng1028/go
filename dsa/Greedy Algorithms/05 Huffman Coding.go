package main

import "fmt"

const MAX_TREE_HT = 50

type MinHNode struct {
	item        rune
	freq        int
	left, right *MinHNode
}
type MinH struct {
	size     int
	capacity int
	array    []*MinHNode
}

// Create Huffman tree node
func newNode(item rune, freq int) *MinHNode {
	return &MinHNode{
		item: item,
		freq: freq,
	}
}

// Create min heap using given capacity
func createMinH(capacity int) *MinH {
	minHeap := &MinH{}
	minHeap.size = 0
	minHeap.capacity = capacity
	minHeap.array = make([]*MinHNode, capacity)
	return minHeap
}

// Print the array
func printArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
}

func swapMinHNode(a, b *MinHNode) {
	temp := *a
	*a = *b
	*b = temp
}

// Heapify
func minHeapify(minHeap *MinH, idx int) {
	smallest := idx
	left := 2*idx + 1
	right := 2*idx + 2

	if left < minHeap.size && minHeap.array[left].freq < minHeap.array[smallest].freq {
		smallest = left
	}

	if right < minHeap.size && minHeap.array[right].freq < minHeap.array[smallest].freq {
		smallest = right
	}

	if smallest != idx {
		swapMinHNode(minHeap.array[smallest], minHeap.array[idx])
		minHeapify(minHeap, smallest)
	}
}

// Check if size is 1
func checkSizeOne(minHeap *MinH) bool {
	return minHeap.size == 1
}

// Extract the min
func extractMin(minHeap *MinH) *MinHNode {
	temp := minHeap.array[0]
	minHeap.array[0] = minHeap.array[minHeap.size-1]

	minHeap.size--
	minHeapify(minHeap, 0)

	return temp
}

// Insertion
func insertMinHeap(minHeap *MinH, minHeapNode *MinHNode) {
	minHeap.size++
	minHeap.array[minHeap.size-1] = minHeapNode
	buildMinHeap(minHeap)
}

// Build min heap
func buildMinHeap(minHeap *MinH) {
	for i := minHeap.size/2 - 1; i >= 0; i-- {
		minHeapify(minHeap, i)
	}
}

func isLeaf(root *MinHNode) bool {
	return root.left == nil && root.right == nil
}

func createAndBuildMinHeap(item []rune, freq []int, size int) *MinH {
	minHeap := createMinH(size)

	for i := 0; i < size; i++ {
		minHeap.array[i] = newNode(item[i], freq[i])
	}

	minHeap.size = size
	buildMinHeap(minHeap)

	return minHeap
}

func buildHfTree(item []rune, freq []int, size int) *MinHNode {
	minHeap := createAndBuildMinHeap(item, freq, size)

	for !checkSizeOne(minHeap) {
		left := extractMin(minHeap)
		right := extractMin(minHeap)

		top := newNode('$', left.freq+right.freq)

		top.left = left
		top.right = right

		insertMinHeap(minHeap, top)
	}

	return extractMin(minHeap)
}

func printHCodes(root *MinHNode, arr []int, top int) {
	if root.left != nil {
		arr[top] = 0
		printHCodes(root.left, arr, top+1)
	}

	if root.right != nil {
		arr[top] = 1
		printHCodes(root.right, arr, top+1)
	}

	if isLeaf(root) {
		fmt.Printf("%c    | ", root.item)
		printArray(arr, top)
	}
}

func HuffmanCodes(item []rune, freq []int, size int) {
	root := buildHfTree(item, freq, size)

	arr := make([]int, MAX_TREE_HT)
	top := 0

	printHCodes(root, arr, top)
}

func main() {
	arr := []rune{'A', 'B', 'C', 'D'}
	freq := []int{5, 1, 6, 3}

	fmt.Println("Char | Huffman code")
	fmt.Println("-------------------")
	HuffmanCodes(arr, freq, len(arr))
}
