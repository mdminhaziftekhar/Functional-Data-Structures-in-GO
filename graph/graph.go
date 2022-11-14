package main

import (
	"fmt"
)

// Graph represent an adjacency list graph
type Graph struct {
	vertices []vertex
}

// Vertex represent a graph
type vertex struct {
	key      int
	adjacent []vertex
}

// Add Edge
func (g Graph) addEdge(from, to int) Graph {
	//Get vertex
	fromVertex, err1 := g.getVertex(from)
	toVertex, err2 := g.getVertex(to)
	// Check error
	if err1 == false || err2 == false {
		err := fmt.Errorf("Invalid edge () %v --> %v )", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) && err1 == true {
		err := fmt.Errorf("Already  edge () %v --> %v )", from, to)
		fmt.Println(err.Error())
	} else {
		// Add edge

		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	}
	g.vertices = append(g.vertices, fromVertex.adjacent...)

	return g
}

// Get vertex
func (g Graph) getVertex(k int) (vertex, bool) {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i], true
		}
	}

	return g.vertices[0], false
}

// addvertex adds a Vertex to the Graph
func (g Graph) addVertex(k int) Graph {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an exisring key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, vertex{key: k})
	}

	return g
}

// printing the graph
func (g Graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v: ", v.key)
		for _, u := range v.adjacent {
			fmt.Printf("%v ", u.key)
		}
	}
}

// vertiecs contaion or not check
func contains(s []vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
func (g Graph) valSet(m, n int) Graph {

	if m == n {
		c := Graph{}
		return c
	}
	g.valSet(m, n-1)
	g.addVertex(n)

	return g
}
func main() {
	test := Graph{}
	for i := 0; i < 5; i++ {
		test = test.addVertex(i)
	}
	//n := 5
	//test = test.valSet(0, n)
	test = test.addVertex(0)
	test = test.addVertex(4)
	test = test.addVertex(1)
	test = test.addVertex(2)
	test = test.addEdge(1, 2)
	test = test.addEdge(2, 3)
	test = test.addEdge(0, 2)
	test = test.addEdge(2, 4)
	test = test.addEdge(3, 3)
	test = test.addEdge(6, 2)
	test = test.addEdge(2, 4)
	test.print()
}
