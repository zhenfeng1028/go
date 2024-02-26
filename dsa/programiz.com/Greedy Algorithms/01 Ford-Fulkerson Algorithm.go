package main

import (
	"fmt"
	"math"
)

const V = 6 // number of vertices

// Use BFS as a searching algorithm
func bfs(rGraph *[V][V]int, s, t int, parent *[V]int) bool {
	visited := make([]bool, V)
	queue := make([]int, 0)
	queue = append(queue, s)
	visited[s] = true
	parent[s] = -1

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]
		for v := 0; v < V; v++ {
			if !visited[v] && rGraph[u][v] > 0 {
				queue = append(queue, v)
				parent[v] = u
				visited[v] = true
			}
		}
	}

	return visited[t]
}

func fordFulkerson(graph [V][V]int, s, t int) int {
	var u, v int

	var rGraph [V][V]int
	for u = 0; u < V; u++ {
		for v = 0; v < V; v++ {
			rGraph[u][v] = graph[u][v]
		}
	}

	var parent [V]int
	var max_flow int

	// Update the residual values of edges
	for bfs(&rGraph, s, t, &parent) {
		path_flow := math.MaxInt
		for v = t; v != s; v = parent[v] {
			u = parent[v]
			path_flow = min(path_flow, rGraph[u][v])
		}

		for v = t; v != s; v = parent[v] {
			u = parent[v]
			rGraph[u][v] -= path_flow
		}

		// Add the path flows
		max_flow += path_flow
	}

	return max_flow
}

func main() {
	graph := [V][V]int{{0, 8, 0, 0, 3, 0}, {0, 0, 9, 0, 0, 0}, {0, 0, 0, 0, 7, 2}, {0, 0, 0, 0, 0, 5}, {0, 0, 7, 4, 0, 0}, {0, 0, 0, 0, 0, 0}}

	fmt.Println("Max Flow:", fordFulkerson(graph, 0, 5))
}
