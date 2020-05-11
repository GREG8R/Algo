package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_Run(t *testing.T) {
	array := [][]Node{
		{{1, 7}, {2, 9}, {5, 14}},
		{{0, 7}, {2, 10}, {3, 15}},
		{{0, 9}, {1, 10}, {3, 11}, {5, 2}},
		{{1, 15}, {2, 11}, {4, 6}},
		{{3, 6}, {5, 9}},
		{{4, 9}, {2, 2}, {0, 14}},
	}

	graph := Build(6, 4, array)

	t.Run("test Dijkstra", func(t *testing.T) {

		compDistance := []int{0, 7, 9, 20, 20, 11}

		comp := []Edge{
			{0, 1},
			{0, 2},
			{2, 3},
			{5, 4},
			{2, 5},
		}

		edges, distance := graph.Dijkstra(0)

		for i, val := range distance {
			assert.Equal(t, val, compDistance[i])
		}

		for i, v := range edges {
			assert.Equal(t, comp[i].v1, v.v1)
			assert.Equal(t, comp[i].v2, v.v2)
		}
	})
}
