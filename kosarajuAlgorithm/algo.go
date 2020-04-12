package main

type GraphAdjacencyVector struct {
	adjacencyVector   [][]int
	countOfNodes      int
	capacityOfVectors int
}

func Build(countOfNodes, capacityOfVectors int) *GraphAdjacencyVector {
	graph := &GraphAdjacencyVector{
		adjacencyVector:   make([][]int, countOfNodes),
		countOfNodes:      countOfNodes,
		capacityOfVectors: capacityOfVectors,
	}

	for i := 0; i < countOfNodes; i++ {
		graph.adjacencyVector[i] = make([]int, 0, capacityOfVectors)
	}

	return graph
}

func (graph *GraphAdjacencyVector) GetStronglyConnectedComponents() []int {
	invertedGraph := graph.invert()
	visited := make([]bool, graph.countOfNodes)
	queue := make([]int, 0, graph.countOfNodes)
	components := make([]int, graph.countOfNodes)

	for i := range invertedGraph.adjacencyVector {
		if !visited[i] {
			invertedGraph.GFS1(visited, &queue, i)
		}
	}

	componentIndex := 1
	for i := len(queue) - 1; i >= 0; i-- {
		if components[queue[i]] == 0 {
			graph.GFS2(componentIndex, components, queue[i])
			componentIndex++
		}
	}
	return components
}

func (graph *GraphAdjacencyVector) GFS1(visited []bool, queue *[]int, node int) {
	visited[node] = true
	for _, g_i_j := range graph.adjacencyVector[node] {
		if !visited[g_i_j] {
			graph.GFS1(visited, queue, g_i_j)
		}
	}
	*queue = append(*queue, node)
}

func (graph *GraphAdjacencyVector) GFS2(indexComponent int, components []int, node int) {
	components[node] = indexComponent
	for _, g_i_j := range graph.adjacencyVector[node] {
		if components[g_i_j] == 0 {
			graph.GFS2(indexComponent, components, g_i_j)
		}
	}
}

func (graph *GraphAdjacencyVector) invert() *GraphAdjacencyVector {
	invertedGraph := Build(graph.countOfNodes, graph.capacityOfVectors)

	for i := range graph.adjacencyVector {
		for _, g_i_j := range graph.adjacencyVector[i] {
			invertedGraph.adjacencyVector[g_i_j] = append(invertedGraph.adjacencyVector[g_i_j], i)
		}
	}

	return invertedGraph
}
