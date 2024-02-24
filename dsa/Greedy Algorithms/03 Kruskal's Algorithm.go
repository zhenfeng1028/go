package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	u int
	v int
	w int // weight
}

type Graph struct {
	edges  []Edge
	mst    []Edge // minimum spanning tree
	parent []int  // which set vertex belongs to
	V      int    // number of vertices in graph
}

func NewGraph(V int) *Graph {
	g := &Graph{V: V}
	g.edges = make([]Edge, 0)
	g.mst = make([]Edge, 0)
	g.parent = make([]int, V)
	for i := 0; i < V; i++ {
		g.parent[i] = i // initialize each node as independent set
	}
	return g
}

func (g *Graph) AddWeightedEdge(u, v, w int) {
	edge := Edge{u, v, w}
	g.edges = append(g.edges, edge)
}

// Find which set node i belongs to
func (g *Graph) find_set(i int) int {
	// If i is the parent of itself
	if i == g.parent[i] {
		return i
	} else {
		// Else if i is not the parent of itself
		// Then i is not the representation of his set,
		// so we recursively call Find on its parent
		return g.find_set(g.parent[i])
	}
}

// Make node u and node v belong to the same set
func (g *Graph) union_set(u, v int) {
	g.parent[u] = g.parent[v]
}

func (g *Graph) kruskal() {
	// Sort all the edges from low weight to high
	sort.SliceStable(g.edges, func(i, j int) bool {
		return g.edges[i].w < g.edges[j].w
	})
	for i := 0; i < len(g.edges); i++ {
		uRep := g.find_set(g.edges[i].u)
		vRep := g.find_set(g.edges[i].v)
		if uRep != vRep {
			g.mst = append(g.mst, g.edges[i])
			g.union_set(uRep, vRep)
		}
	}
}

func (g *Graph) print() {
	fmt.Println("Edge : Weight")
	for i := 0; i < len(g.mst); i++ {
		fmt.Printf("%d - %d : %d\n", g.mst[i].u, g.mst[i].v, g.mst[i].w)
	}
}

func main() {
	g := NewGraph(6)

	g.AddWeightedEdge(0, 1, 4)
	g.AddWeightedEdge(0, 2, 4)
	g.AddWeightedEdge(1, 2, 2)
	g.AddWeightedEdge(2, 3, 3)
	g.AddWeightedEdge(2, 4, 2)
	g.AddWeightedEdge(2, 5, 4)
	g.AddWeightedEdge(3, 5, 3)
	g.AddWeightedEdge(4, 5, 3)

	g.kruskal()
	g.print()
}
