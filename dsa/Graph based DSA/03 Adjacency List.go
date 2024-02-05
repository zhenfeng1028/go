package main

import (
	"container/list"
	"fmt"
)

type Graph []*list.List

func NewGraph(numVertices int) *Graph {
	g := make(Graph, numVertices)
	for i := range g {
		g[i] = list.New()
	}
	return &g
}

func (g *Graph) addEdge(s, d int) {
	(*g)[s].PushBack(d)
	(*g)[d].PushBack(s)
}

func (g *Graph) printGraph() {
	for i, v := range *g {
		fmt.Printf("Vertex %d", i)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Print(" -> ", e.Value)
		}
		fmt.Println()
	}
}

func main() {
	g := NewGraph(4)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 3)
	g.addEdge(1, 2)

	g.printGraph()
}
