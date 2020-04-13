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

func (graph *GraphAdjacencyVector) TopologySort() [][]int {
	levels := make([][]int, 1)
	level := 0
	levels[level] = make([]int, 0)

	arrayOfSum := make([]int, graph.countOfNodes)
	for _, g_i := range graph.adjacencyVector {
		for _, g_i_j := range g_i {
			arrayOfSum[g_i_j]++
		}
	}

	matrix := graph.makeMatrix()

	i := 0
	for i < graph.countOfNodes {
		zero := make([]int, 0)
		for j, v := range arrayOfSum {
			if v == 0 {
				arrayOfSum[j] = -1
				zero = append(zero, j)
			}
		}

		if len(zero) == 0 {
			panic("zero is empty")
		}

		for _, v := range zero {
			levels[level] = append(levels[level], v)
			i++
			recalculate(matrix[v], arrayOfSum)
		}

		level++

		if i < graph.countOfNodes {
			levels = append(levels, make([]int, 0))
		}
	}

	return levels
}

func recalculate(m []int, array []int) {
	for i, v := range m {
		array[i] -= v
	}
}

func (graph *GraphAdjacencyVector) makeMatrix() [][]int {
	matrix := make([][]int, graph.countOfNodes)
	for i, g_i := range graph.adjacencyVector {
		matrix[i] = make([]int, graph.countOfNodes)
		for _, g_i_j := range g_i {
			matrix[i][g_i_j] = 1
		}
	}
	return matrix
}

//func (graph *GraphAdjacencyVector) GetStronglyConnectedComponents() []int {
//	invertedGraph := graph.invert()
//	visited := make([]bool, graph.countOfNodes)
//	queue := make([]int, 0, graph.countOfNodes)
//	components := make([]int, graph.countOfNodes)
//
//	for i := range invertedGraph.adjacencyVector {
//		if !visited[i] {
//			invertedGraph.GFS1(visited, &queue, i)
//		}
//	}
//
//	componentIndex := 1
//	for i := len(queue) - 1; i >= 0; i-- {
//		if components[queue[i]] == 0 {
//			graph.GFS2(componentIndex, components, queue[i])
//			componentIndex++
//		}
//	}
//	return components
//}

//func (graph *GraphAdjacencyVector) GFS1(visited []bool, queue *[]int, node int) {
//	visited[node] = true
//	for _, g_i_j := range graph.adjacencyVector[node] {
//		if !visited[g_i_j] {
//			graph.GFS1(visited, queue, g_i_j)
//		}
//	}
//	*queue = append(*queue, node)
//}
//
//func (graph *GraphAdjacencyVector) GFS2(indexComponent int, components []int, node int) {
//	components[node] = indexComponent
//	for _, g_i_j := range graph.adjacencyVector[node] {
//		if components[g_i_j] == 0 {
//			graph.GFS2(indexComponent, components, g_i_j)
//		}
//	}
//}
//
//func (graph *GraphAdjacencyVector) invert() *GraphAdjacencyVector {
//	invertedGraph := Build(graph.countOfNodes, graph.capacityOfVectors)
//
//	for i := range graph.adjacencyVector {
//		for _, g_i_j := range graph.adjacencyVector[i] {
//			invertedGraph.adjacencyVector[g_i_j] = append(invertedGraph.adjacencyVector[g_i_j], i)
//		}
//	}
//
//	return invertedGraph
//}
