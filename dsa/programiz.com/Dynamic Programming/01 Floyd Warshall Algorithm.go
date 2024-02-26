package main

import (
	"fmt"
)

const (
	V   = 4   // number of vertices
	INF = 999 // representation of infinity
)

func floydWarshall(graph [V][V]int) {
	matrix := [V][V]int{}

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			matrix[i][j] = graph[i][j]
		}
	}

	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				if matrix[i][k]+matrix[k][j] < matrix[i][j] {
					matrix[i][j] = matrix[i][k] + matrix[k][j]
				}
			}
		}
	}

	printMatrix(matrix)
}

func printMatrix(matrix [V][V]int) {
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if matrix[i][j] == INF {
				fmt.Printf("%-4s", "INF")
			} else {
				fmt.Printf("%-4d", matrix[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	graph := [V][V]int{{0, 3, INF, 5}, {2, 0, INF, 4}, {INF, 1, 0, INF}, {INF, INF, 2, 0}}
	fmt.Println(graph)
	floydWarshall(graph)
}
