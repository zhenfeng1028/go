package main

import (
	"fmt"
	"math"
)

type Node struct {
	id                rune
	previous          *Node
	distanceFromStart int
}

type Edge struct {
	node1    *Node
	node2    *Node
	distance int
}

var (
	nodes = make([]*Node, 0)
	edges = make([]*Edge, 0)
)

func NewNode(id rune) *Node {
	node := &Node{
		id:                id,
		previous:          nil,
		distanceFromStart: math.MaxInt,
	}
	nodes = append(nodes, node)
	return node
}

func NewEdge(node1, node2 *Node, distance int) *Edge {
	edge := &Edge{
		node1:    node1,
		node2:    node2,
		distance: distance,
	}
	edges = append(edges, edge)
	return edge
}

func (e *Edge) Connects(node1, node2 *Node) bool {
	return node1 == e.node1 && node2 == e.node2 || node1 == e.node2 && node2 == e.node1
}

// Find the node with the smallest distance,
// remove it, and return it.
func ExtractSmallest() *Node {
	size := len(nodes)
	if size == 0 {
		return nil
	}
	smallestPosition := 0
	smallest := nodes[0]
	for i := 1; i < size; i++ {
		current := nodes[i]
		if current.distanceFromStart < smallest.distanceFromStart {
			smallest = current
			smallestPosition = i
		}
	}
	nodes = append(nodes[:smallestPosition], nodes[smallestPosition+1:]...)
	return smallest
}

// Return all nodes adjacent to 'node' which are still
// in the 'nodes' collection
func AdjacentRemainingNodes(node *Node) []*Node {
	adjacentNodes := make([]*Node, 0)
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		var adjacent *Node
		if edge.node1 == node {
			adjacent = edge.node2
		} else if edge.node2 == node {
			adjacent = edge.node1
		}
		if adjacent != nil && Contains(adjacent) {
			adjacentNodes = append(adjacentNodes, adjacent)
		}
	}
	return adjacentNodes
}

// Does the 'nodes' contain 'node'
func Contains(node *Node) bool {
	for i := 0; i < len(nodes); i++ {
		if node == nodes[i] {
			return true
		}
	}
	return false
}

// Return distance between two connected nodes
func Distance(node1, node2 *Node) int {
	for i := 0; i < len(nodes); i++ {
		edge := edges[i]
		if edge.Connects(node1, node2) {
			return edge.distance
		}
	}
	return -1 // should never happen
}

func Dijkstras() {
	for len(nodes) > 0 {
		smallest := ExtractSmallest()
		adjacentNodes := AdjacentRemainingNodes(smallest)

		for i := 0; i < len(adjacentNodes); i++ {
			adjacent := adjacentNodes[i]
			distance := Distance(smallest, adjacent) + smallest.distanceFromStart

			if distance < adjacent.distanceFromStart {
				adjacent.distanceFromStart = distance
				adjacent.previous = smallest
			}
		}
	}
}

func PrintShortestRouteTo(destination *Node) {
	previous := destination
	fmt.Println("Distance from start:", destination.distanceFromStart)
	for previous != nil {
		fmt.Printf("%c ", previous.id)
		previous = previous.previous
	}
	fmt.Println()
}

func main() {
	a := NewNode('a')
	b := NewNode('b')
	c := NewNode('c')
	d := NewNode('d')
	e := NewNode('e')
	f := NewNode('f')

	NewEdge(a, b, 4)
	NewEdge(a, c, 4)
	NewEdge(b, c, 2)
	NewEdge(c, d, 3)
	NewEdge(c, e, 1)
	NewEdge(c, f, 6)
	NewEdge(d, f, 2)
	NewEdge(e, f, 3)

	a.distanceFromStart = 0 // set start node
	Dijkstras()
	PrintShortestRouteTo(f)
}
