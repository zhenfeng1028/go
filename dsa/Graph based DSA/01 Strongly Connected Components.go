package main

import "fmt"

type Graph struct {
	V   int     // number of vertices
	adj [][]int // adjacency lists
}

func NewGraph(V int) *Graph {
	return &Graph{
		V:   V,
		adj: make([][]int, V),
	}
}

// DFS
func (g *Graph) DFS(s int, visitedV []bool) {
	visitedV[s] = true
	fmt.Print(s, " ")

	list := g.adj[s]
	for i := 0; i < len(list); i++ {
		if !visitedV[list[i]] {
			g.DFS(list[i], visitedV)
		}
	}
}

// Transpose
func (g *Graph) Transpose() *Graph {
	gr := NewGraph(g.V)
	for s := 0; s < g.V; s++ {
		list := g.adj[s]
		for i := 0; i < len(list); i++ {
			gr.adj[list[i]] = append(gr.adj[list[i]], s)
		}
	}
	return gr
}

// Add edge into the graph
func (g *Graph) addEdge(s, d int) {
	g.adj[s] = append(g.adj[s], d)
}

func (g *Graph) fillOrder(s int, visitedV []bool, stack *[]int) {
	visitedV[s] = true
	list := g.adj[s]
	for i := 0; i < len(list); i++ {
		if !visitedV[list[i]] {
			g.fillOrder(list[i], visitedV, stack)
		}
	}
	*stack = append(*stack, s)
}

func (g *Graph) printAdj() {
	for s := 0; s < g.V; s++ {
		fmt.Print(s, ": ")
		list := g.adj[s]
		for i := 0; i < len(list); i++ {
			fmt.Print(list[i], " ")
		}
		fmt.Println()
	}
}

// Print strongly connected component
func (g *Graph) printSCC() {
	stack := make([]int, 0)
	visitedV := make([]bool, g.V)

	for i := 0; i < g.V; i++ {
		if !visitedV[i] {
			g.fillOrder(i, visitedV, &stack)
		}
	}

	gr := g.Transpose()

	// reset visitedV to all false
	for i := 0; i < g.V; i++ {
		visitedV[i] = false
	}

	for len(stack) != 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visitedV[s] {
			gr.DFS(s, visitedV)
			fmt.Println()
		}
	}
}

func main() {
	g := NewGraph(8)

	g.addEdge(0, 1)
	g.addEdge(1, 2)
	g.addEdge(2, 3)
	g.addEdge(2, 4)
	g.addEdge(3, 0)
	g.addEdge(4, 5)
	g.addEdge(5, 6)
	g.addEdge(6, 4)
	g.addEdge(6, 7)

	g.printAdj()

	fmt.Println("Strongly Connected Components:")

	g.printSCC()
}
