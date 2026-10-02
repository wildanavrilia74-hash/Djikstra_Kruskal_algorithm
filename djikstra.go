package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

func dijkstra(graph [][]Edge, start int) []int {
	n := len(graph)

	// Inisialisasi jarak dengan nilai tak hingga
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}

	dist[start] = 0

	for i := 0; i < n; i++ {
		// Cari vertex dengan jarak terkecil
		u := -1
		for j := 0; j < n; j++ {
			if !visited[j] && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}

		if u == -1 || dist[u] == math.MaxInt {
			break
		}

		visited[u] = true

		// Update jarak vertex tetangga
		for _, edge := range graph[u] {
			v := edge.to
			newDist := dist[u] + edge.weight

			if newDist < dist[v] {
				dist[v] = newDist
			}
		}
	}

	return dist
}

func main() {
	// Membuat graph
	graph := make([][]Edge, 5)

	graph[0] = append(graph[0], Edge{1, 4})
	graph[0] = append(graph[0], Edge{2, 1})

	graph[1] = append(graph[1], Edge{0, 4})
	graph[1] = append(graph[1], Edge{2, 2})
	graph[1] = append(graph[1], Edge{3, 5})

	graph[2] = append(graph[2], Edge{0, 1})
	graph[2] = append(graph[2], Edge{1, 2})
	graph[2] = append(graph[2], Edge{3, 8})
	graph[2] = append(graph[2], Edge{4, 10})

	graph[3] = append(graph[3], Edge{1, 5})
	graph[3] = append(graph[3], Edge{2, 8})
	graph[3] = append(graph[3], Edge{4, 2})

	graph[4] = append(graph[4], Edge{2, 10})
	graph[4] = append(graph[4], Edge{3, 2})

	// Mulai dari vertex 0
	dist := dijkstra(graph, 0)

	fmt.Println("Jarak terpendek dari vertex 0:")

	for i, d := range dist {
		fmt.Printf("0 -> %d = %d\n", i, d)
	}
}
