package main

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRedBlackTree_Run(t *testing.T) {

	t.Run("test insert search and treeValues", func(t *testing.T) {
		array := []int{4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0, 7}
		sortArray := []int{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9}

		tree := Build(array)

		var arr []int
		tree.Root.Values(&arr)
		assert.Equal(t, arr, sortArray)

		assert.Equal(t, tree.Root.Search(9), true)
		assert.Equal(t, tree.Root.Search(0), true)
		assert.Equal(t, tree.Root.Search(2), true)
		assert.Equal(t, tree.Root.Search(4), true)
		assert.Equal(t, tree.Root.Search(6), true)
		assert.Equal(t, tree.Root.Search(10), false)
	})

	t.Run("test tree", func(t *testing.T) {
		sortArray := []int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 10, 10, 10, 11, 12, 13, 15}
		tree := RedBlackTree{
			Root: nil,
		}

		for _, v := range sortArray {
			tree.Insert(v)
		}
		fmt.Println()
	})
}
