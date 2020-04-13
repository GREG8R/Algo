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
