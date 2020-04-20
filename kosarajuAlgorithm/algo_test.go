package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_Run(t *testing.T) {
	array := [][]int{
		{1},
		{2, 4, 5},
		{3, 6},
		{2, 7},
		{0, 5},
		{6},
		{5},
		{3, 6},
	}

	graph := Build(8, 3)
	for i, g_i := range array {
		for _, g_i_j := range g_i {
			graph.adjacencyVector[i] = append(graph.adjacencyVector[i], g_i_j)
		}
	}

	t.Run("test invert", func(t *testing.T) {
		invertGraph := graph.invert()

		invertArray := [][]int{
			{4},
			{0},
			{1, 3},
			{2, 7},
			{1},
			{1, 4, 6},
			{2, 5, 7},
			{3},
		}

		for i, g_i := range invertArray {
			for j, g_i_j := range g_i {
				assert.Equal(t, invertGraph.adjacencyVector[i][j], g_i_j)
			}
		}
	})

	t.Run("test kosaraju", func(t *testing.T) {
		comp := []int{3, 3, 2, 2, 3, 1, 1, 2}

		result := graph.GetStronglyConnectedComponents()

		for i, v := range result {
			assert.Equal(t, comp[i], v)
		}
	})
}
