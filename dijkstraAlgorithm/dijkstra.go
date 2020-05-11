package main

import (
	"math"
)

type Edge struct {
	v1, v2 int
}

type Node struct {
	v     int
	value int
}

type Point struct {
	v         int
	isDeleted bool
}

type GraphAdjacencyVector struct {
	adjacencyVector   [][]Node
	countOfNodes      int
	capacityOfVectors int
}

func (graph *GraphAdjacencyVector) Dijkstra(a int) ([]Edge, []int) {

	points := make([]Point, graph.countOfNodes)
	path := make([]int, graph.countOfNodes)
	distance := make([]int, graph.countOfNodes)

	for i := range points {
		points[i] = Point{
			v:         math.MaxInt64,
			isDeleted: false,
		}
	}
	points[a].v = 0

	for {
		if isEmpty(points) {
			break
		}

		v := minValue(points)

		for _, e := range graph.adjacencyVector[v] {
			if contains(points, e.v) {
				if points[v].v+e.value < points[e.v].v {
					points[e.v].v = points[v].v + e.value
					path[e.v] = v
				}
			}
		}
		distance[v] = points[v].v
		points[v].isDeleted = true
	}

	result := make([]Edge, 0, graph.countOfNodes)

	for i, val := range path {
		if i != a {
			result = append(result, Edge{
				v1: val,
				v2: i,
			})
		}
	}

	return result, distance
}

func minValue(p []Point) int {
	minV := math.MaxInt32
	index := -1
	for i, v := range p {
		if v.v < minV && !v.isDeleted {
			minV = v.v
			index = i
		}
	}
	return index
}

func contains(p []Point, v int) bool {
	for i, val := range p {
		if i == v && !val.isDeleted {
			return true
		}
	}
	return false
}

func isEmpty(points []Point) bool {
	for _, val := range points {
		if !val.isDeleted {
			return false
		}
	}
	return true
}

func Build(countOfNodes, capacityOfVectors int, array [][]Node) *GraphAdjacencyVector {
	graph := &GraphAdjacencyVector{
		adjacencyVector:   make([][]Node, countOfNodes),
		countOfNodes:      countOfNodes,
		capacityOfVectors: capacityOfVectors,
	}

	for i := 0; i < countOfNodes; i++ {
		graph.adjacencyVector[i] = make([]Node, 0, capacityOfVectors)
	}

	for i, g_i := range array {
		for _, g_i_j := range g_i {
			graph.adjacencyVector[i] = append(graph.adjacencyVector[i], g_i_j)
		}
	}

	return graph
}
