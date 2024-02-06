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

func (g *Graph) BFS(startVertex int) {
	queue := []int{}
	g.visited[startVertex] = true
	queue = append(queue, startVertex)
	for len(queue) != 0 {
		currentVertex := queue[len(queue)-1]
		fmt.Print(currentVertex, " ")
		queue = queue[:len(queue)-1]
		adjList := g.adjLists[currentVertex]
		for e := adjList.Front(); e != nil; e = e.Next() {
			if adjVertex, ok := e.Value.(int); ok {
				if !g.visited[adjVertex] {
					g.visited[adjVertex] = true
					queue = append(queue, adjVertex)
				}
			}
		}
	}
}

func main() {
	g := NewGraph(4)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	g.BFS(2)
}
