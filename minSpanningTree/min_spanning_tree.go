package main

import (
	"sort"
)

type Edge struct {
	v1, v2 int
	value  int
}

type Node struct {
	v     int
	value int
}

type GraphAdjacencyVector struct {
	adjacencyVector   [][]Node
	countOfNodes      int
	capacityOfVectors int
}

func Build(countOfNodes, capacityOfVectors int) *GraphAdjacencyVector {
	graph := &GraphAdjacencyVector{
		adjacencyVector:   make([][]Node, countOfNodes),
		countOfNodes:      countOfNodes,
		capacityOfVectors: capacityOfVectors,
	}

	for i := 0; i < countOfNodes; i++ {
		graph.adjacencyVector[i] = make([]Node, 0, capacityOfVectors)
	}

	return graph
}

func (g *GraphAdjacencyVector) KruskalAlgo() []Edge {
	result := make([]Edge, 0)

	// init array with nodes index and parent values
	nodesWithParent := make([]int, g.countOfNodes)
	for i := 0; i < g.countOfNodes; i++ {
		nodesWithParent[i] = i
	}

	allEdges := g.buildListOfEdges()

	for _, v := range allEdges {
		if nodesWithParent[v.v1] != nodesWithParent[v.v2] {
			merge(nodesWithParent, v)
			result = append(result, v)
		}
	}

	return result
}

func (g *GraphAdjacencyVector) buildListOfEdges() []Edge {
	result := make([]Edge, 0)

	for i, g_i := range g.adjacencyVector {
		for _, g_i_j := range g_i {
			result = append(result, Edge{
				v1:    i,
				v2:    g_i_j.v,
				value: g_i_j.value,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].value < result[j].value
	})

	return result
}

func merge(nodesWithParent []int, v Edge) {
	parentValueV2 := nodesWithParent[v.v2]
	nodesWithParent[v.v2] = nodesWithParent[v.v1]
	for i := range nodesWithParent {
		if nodesWithParent[i] == parentValueV2 {
			nodesWithParent[i] = nodesWithParent[v.v1]
		}
	}
}
