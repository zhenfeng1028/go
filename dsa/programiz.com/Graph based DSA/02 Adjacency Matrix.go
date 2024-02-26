package main

import "fmt"

type Graph struct {
	numVertices int
	adjMatrix   [][]int
}

func NewGraph(numVertices int) *Graph {
	g := &Graph{}
	g.numVertices = numVertices
	g.adjMatrix = make([][]int, numVertices)
	for i := range g.adjMatrix {
		g.adjMatrix[i] = make([]int, numVertices)
	}
	return g
}

func (g *Graph) addEdge(i, j int) {
	g.adjMatrix[i][j] = 1
	g.adjMatrix[j][i] = 1
}

func (g *Graph) removeEdge(i, j int) {
	g.adjMatrix[i][j] = 0
	g.adjMatrix[j][i] = 0
}

func (g *Graph) printMatrix() {
	for i := 0; i < g.numVertices; i++ {
		for j := 0; j < g.numVertices; j++ {
			fmt.Print(g.adjMatrix[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	g := NewGraph(4)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)

	g.printMatrix()
}
