package main

import (
	"fmt"
)

func PrintPath(path [][]float64) {
	for i := 0; i < len(path); i++ {
		for j := 0; j < len(path[i]); j++ {
			fmt.Printf("%.2f ", path[i][j])
		}
		fmt.Println()
	}
}

/*
Функція main є вхідним пунктом в програму і містить код для
створення та взаємодії з .
*/
func main() {
	// Тестування піраміди Фібоначі
	heap := NewFibHeap()
	heap.Insert(1, 4)
	heap.Insert(2, 8)
	heap.Insert(3, 2)
	heap.Insert(4, 45)
	heap.Insert(5, 3)
	heap.Insert(453534, 1)
	heap.ExtractMin()

	print(heap.String())

	// Створення графа
	g := Graph{}

	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 3, 1)
	g.AddEdge(1, 2, 0)
	g.AddEdge(1, 0, 2)
	g.AddEdge(2, 0, 3)
	g.AddEdge(3, 0, 0)
	g.AddEdge(3, 2, 5)

	fmt.Println("BellmanFord (from 0 vertex):")
	src := 0
	dist, prev, hasNeg := g.BellmanFord(src)

	fmt.Println(hasNeg)
	fmt.Printf("Shortest paths from vertex %d:\n", src)
	for i := 0; i < len(g.vertices); i++ {
		fmt.Printf("Vertex %d: Distance: %f, Previous: %d\n", g.vertices[i], dist[i], prev[i])
	}

	fmt.Println("Johnson:")
	PrintPath(g.Johnson())
}
