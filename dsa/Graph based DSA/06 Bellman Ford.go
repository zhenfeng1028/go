package main

import (
	"fmt"
	"math"
)

type Edge struct {
	u int // start vertex of the edge
	v int // end vertex of the edge
	w int // weight of the edge (u,v)
}

type Graph struct {
	V     int     // total number of vertices in the graph
	E     int     // total number of edges in the graph
	edges []*Edge // array of edges
}

func createGraph(V, E int) *Graph {
	g := &Graph{}
	g.V = V
	g.E = E
	g.edges = make([]*Edge, E)
	for i := 0; i < E; i++ {
		g.edges[i] = &Edge{}
	}
	return g
}

func BellmanFord(g *Graph, src int) {
	V := g.V
	E := g.E
	dist := make([]int, V)

	// Initialize distance of all vertices as infinite
	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt
	}

	// initialize distance of source as 0
	dist[src] = 0

	// Relax all edges |V| - 1 times. A simple
	// shortest path from src to any other
	// vertex can have at-most |V| - 1 edges
	for i := 1; i <= V-1; i++ {
		for j := 0; j < E; j++ {
			u := g.edges[j].u
			v := g.edges[j].v
			w := g.edges[j].w
			if dist[u] != math.MaxInt && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
			}
		}
	}

	// Check for negative-weight cycles.
	// The above step guarantees shortest
	// distances if graph doesn't contain
	// negative weight cycle. If we get a
	// shorter path, then there is a cycle.
	for i := 0; i < E; i++ {
		u := g.edges[i].u
		v := g.edges[i].v
		w := g.edges[i].w
		if dist[u] != math.MaxInt && dist[u]+w < dist[v] {
			fmt.Print("Graph contains negative weight cycle")
			return
		}
	}

	fmt.Println("Vertex distance from source")
	for i := 0; i < V; i++ {
		fmt.Printf("%d\t%d\n", i, dist[i])
	}
}

func main() {
	V, E := 4, 5
	g := createGraph(V, E)

	//------- adding the edges of the graph

	// edge 0 --> 1
	g.edges[0].u = 0
	g.edges[0].v = 1
	g.edges[0].w = 5

	// edge 0 --> 2
	g.edges[1].u = 0
	g.edges[1].v = 2
	g.edges[1].w = 3

	// edge 1 --> 3
	g.edges[2].u = 1
	g.edges[2].v = 3
	g.edges[2].w = 3

	// edge 2 --> 1
	g.edges[3].u = 2
	g.edges[3].v = 1
	g.edges[3].w = 1

	// edge 3 --> 2
	g.edges[4].u = 3
	g.edges[4].v = 2
	g.edges[4].w = 2

	BellmanFord(g, 0)
}
