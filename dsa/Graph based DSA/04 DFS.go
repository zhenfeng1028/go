package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	visited  []bool
	adjLists []*list.List
}

func NewGraph(numVertices int) *Graph {
	g := &Graph{
		visited:  make([]bool, numVertices),
		adjLists: make([]*list.List, numVertices),
	}
	for i := range g.adjLists {
		g.adjLists[i] = list.New()
	}
	return g
}

func (g *Graph) addEdge(s, d int) {
	g.adjLists[s].PushFront(d)
	g.adjLists[d].PushFront(s)
}

func (g *Graph) DFS(vertex int) {
	g.visited[vertex] = true
	adjList := g.adjLists[vertex]

	fmt.Print(vertex, " ")

	for e := adjList.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(int); ok {
			if !g.visited[v] {
				g.DFS(v)
			}
		}
	}
}

func main() {
	g := NewGraph(6)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 3)
	g.addEdge(4, 5) // disconnected part of the graph

	g.DFS(2)
}
