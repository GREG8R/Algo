package main

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRedBlackTree_Run(t *testing.T) {

	t.Run("test Kruskal's algo", func(t *testing.T) {
		array := [][]Node{
			{{1, 7}, {3, 5}},
			{{2, 8}, {3, 9}, {4, 7}},
			{{4, 5}},
			{{4, 15}, {5, 6}},
			{{5, 8}, {6, 9}},
			{{6, 11}},
			{},
		}

		graph := Build(7, 3)
		for i, g_i := range array {
			for _, g_i_j := range g_i {
				graph.adjacencyVector[i] = append(graph.adjacencyVector[i], g_i_j)
			}
		}

		result := graph.KruskalAlgo()

		equalResult := []Edge{
			{0, 3, 5},
			{2, 4, 5},
			{3, 5, 6},
			{0, 1, 7},
			{1, 4, 7},
			{4, 6, 9},
		}

		sum := 0
		for i, v := range result {
			fmt.Printf("edge = from %d to %d, value = %d\n", v.v1, v.v2, v.value)
			assert.Equal(t, equalResult[i], v)
			sum += v.value
		}

		fmt.Printf("sum = %d\n", sum)
		assert.Equal(t, sum, 39)

	})
}
