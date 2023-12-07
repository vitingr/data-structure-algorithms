package main

import (
	"fmt"
	"log"
)

// -=-=-=-=-=-=-=- GRAPHS -=-=-=-=-=-=-=-=-=-
// A graph have nodes and edges (childs), is like a tree, but a crazy tree with so many connections
// When the graph has many address we call "Dense" and the nodes turn into vertex, when has less address we call "Sparse"
// Is a abstracted data structure, we can build with Directed Graph, Undirected Graph, Cyclic Graph, Weighted Graph
// Weighted Graph is like the google maps, you have points and you can calculate the distance between them
// No circles (like a binary tree) are allowed and you cant interact the child with the parent directly (max 2 node children)

// Graph Structure represents an adjacent list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex Structure represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		log.Fatal("Cannot add this vertex key because already exists in this graph")
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// Get Vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// Check Error
	if fromVertex == nil || toVertex == nil {
		log.Fatal("You can't insert a new edge into a Vertex that doesn't exists")
	} else if contains(fromVertex.adjacent, to) {
		log.Fatal("You can't insert this edge, because already exists")
	} else {
		// Add Edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, value := range g.vertices {
		if value.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// Contains
func contains(s []*Vertex, k int) bool {
	// Will search into the entire array to search if already exists this value
	for _, value := range s {
		if k == value.key {
			return true
		}
	}
	return false
}

// Print will show the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, value := range g.vertices {
		fmt.Printf("\nVertex %v: ", value.key)
		for _, value := range value.adjacent {
			fmt.Printf("%v", value.key)
		}
	}
}

func main() {
	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	graph.Print()
}
