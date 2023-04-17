package main

import (
	"math"
)

// Edge represents a weighted edge in a graph
type Edge struct {
	src, dest int
	weight    float64
}

// Graph represents a weighted graph using an adjacency list
type Graph struct {
	edges    []Edge
	vertices []int
}

// AddEdge adds a new weighted edge to the graph
func (g *Graph) AddEdge(src, dest int, weight float64) {
	edge := Edge{src, dest, weight}
	g.edges = append(g.edges, edge)

	var found bool = false
	for i := 0; i < len(g.vertices); i++ {
		if g.vertices[i] == src {
			found = true
			break
		}
	}

	if !found {
		g.vertices = append(g.vertices, src)
	}

	found = false
	for i := 0; i < len(g.vertices); i++ {
		if g.vertices[i] == dest {
			found = true
			break
		}
	}

	if !found {
		g.vertices = append(g.vertices, dest)
	}
}

// BellmanFord detects negative weight cycles and returns the shortest paths
// from the source vertex to all other vertices in the graph
func (g *Graph) BellmanFord(src int) ([]float64, []int, bool) {
	n := len(g.vertices)
	dist := make([]float64, n)
	prev := make([]int, n)

	// Initialize distance array with infinity and previous array with -1
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}

	// Set distance from source vertex to itself as 0
	dist[src] = 0

	// Relax edges n-1 times
	for i := 1; i < n; i++ {
		for _, edge := range g.edges {
			u, v, w := edge.src, edge.dest, edge.weight
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				prev[v] = u
			}
		}
	}

	// Check for negative weight cycles by relaxing edges one more time
	for _, edge := range g.edges {
		u, v, w := edge.src, edge.dest, edge.weight
		if dist[u]+w < dist[v] {
			return nil, nil, false
		}
	}

	return dist, prev, true
}

// Dijkstra returns the shortest paths from the source vertex to all other vertices
// in the graph using Dijkstra's algorithm
func (g *Graph) Dijkstra(src int, pq HeapInterface) ([]float64, []int) {
	n := len(g.vertices)
	dist := make([]float64, n)
	prev := make([]int, n)
	visited := make([]bool, n)

	// Initialize distance array with infinity and previous array with -1
	for i := 0; i < n; i++ {
		dist[i] = math.Inf(1)
		prev[i] = -1
	}

	// Set distance from source vertex to itself as 0
	dist[src] = 0

	// Insert source vertex into priority queue
	pq.Insert(src, float64(dist[src]))

	for pq.Num() > 0 {
		// Extract minimum distance vertex from priority queue
		u, _ := pq.ExtractMin()

		if visited[u.(int)] {
			continue
		}

		visited[u.(int)] = true

		// Relax edges
		for _, edge := range g.edges {
			if edge.src == u.(int) {
				v := edge.dest
				w := edge.weight
				if dist[u.(int)]+w < dist[v] {
					dist[v] = dist[u.(int)] + w
					prev[v] = u.(int)
					pq.Insert(v, float64(dist[v]))
				}
			}
		}
	}

	return dist, prev
}

// Johnson computes the shortest paths between all pairs of vertices in the graph
// using Johnson's algorithm
func (g *Graph) Johnson() [][]float64 {
	n := len(g.vertices)

	// Add a new vertex s to the graph and edges of weight 0 from s to all other vertices
	g.AddEdge(n, n+1, 0)
	for i := 0; i < n; i++ {
		g.AddEdge(n, i, 0)
	}

	// Run Bellman-Ford algorithm to get the potential function for each vertex
	potential, _, hasNegativeCycle := g.BellmanFord(n)

	// If there is a negative weight cycle in the graph, return nil
	if !hasNegativeCycle {
		return nil
	}

	// Remove the added vertex and edges
	g.edges = g.edges[:len(g.edges)-n-1]
	for i := 0; i < n; i++ {
		g.edges = g.edges[:len(g.edges)-1]
	}

	// Create a new graph with edges of weight w(u,v) + h(u) - h(v)
	newGraph := &Graph{vertices: g.vertices}
	for _, edge := range g.edges {
		u, v, w := edge.src, edge.dest, edge.weight
		newW := w + potential[u] - potential[v]
		newGraph.AddEdge(u, v, newW)
	}

	// Initialize a 2D slice to store the shortest paths between all pairs of vertices
	paths := make([][]float64, n)
	for i := 0; i < n; i++ {
		paths[i] = make([]float64, n)
	}

	// Run Dijkstra's algorithm from each vertex to get the shortest paths to all other vertices
	for i := 0; i < n; i++ {
		dist, _ := newGraph.Dijkstra(i, NewFibHeap())

		// Adjust the distances using the potential function to get the actual distances
		for j := 0; j < n; j++ {
			paths[i][j] = dist[j] - potential[i] + potential[j]
		}
	}

	return paths
}
