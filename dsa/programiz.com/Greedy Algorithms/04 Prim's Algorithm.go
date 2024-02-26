package main

import (
	"fmt"
	"math"
)

const V = 5 // number of vertices

var (
	no_edge  int     // number of edges
	selected [V]bool // an array to track selected vertex
)

func prim(graph [V][V]int) {
	// choose 0th vertex and make it true
	selected[0] = true

	fmt.Println("Edge : Weight")

	// the number of egde in minimum spanning tree will be always
	// less than (V -1), where V is number of vertices in graph
	for no_edge < V-1 {
		// For every vertex in the set S, find the all adjacent vertices.
		// If the vertex is already in the set S, discard it and
		// choose another vertex nearest to selected vertex.
		min := math.MaxInt
		x := 0 // row number
		y := 0 // col number

		for i := 0; i < V; i++ {
			if selected[i] {
				for j := 0; j < V; j++ {
					if !selected[j] && graph[i][j] > 0 { // not in selected and there is an edge
						if min > graph[i][j] {
							min = graph[i][j]
							x = i
							y = j
						}
					}
				}
			}
		}

		fmt.Printf("%d - %d : %d\n", x, y, graph[x][y])
		selected[y] = true
		no_edge++
	}
}

func main() {
	graph := [V][V]int{{0, 9, 75, 0, 0}, {9, 0, 95, 19, 42}, {75, 95, 0, 51, 66}, {0, 19, 51, 0, 31}, {0, 42, 66, 31, 0}}

	prim(graph)
}
