package main

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_Run(t *testing.T) {

	t.Run("test topology sort", func(t *testing.T) {
		array := [][]int{
			{2, 12},
			{12},
			{},
			{2},
			{2, 8, 9},
			{3, 10, 11},
			{10},
			{1, 3, 5, 6},
			{0, 13},
			{0, 6, 11},
			{2},
			{},
			{2},
			{5},
		}

		graph := Build(14, 4)
		for i, g_i := range array {
			for _, g_i_j := range g_i {
				graph.adjacencyVector[i] = append(graph.adjacencyVector[i], g_i_j)
			}
		}

		result := graph.TopologySort()

		equalResult := [][]int{
			{4, 7},
			{1, 8, 9},
			{0, 6, 13},
			{5, 12},
			{3, 10, 11},
			{2},
		}

		for i, v := range result {
			str := fmt.Sprintf("level = %d ", i)
			for j, vj := range v {
				str += fmt.Sprintf("node = %d, ", vj)
				assert.Equal(t, vj, equalResult[i][j])
			}
			fmt.Println(str)
		}
	})
}
