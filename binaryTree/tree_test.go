package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInsertBinaryTree_Run(t *testing.T) {
	t.Run("test insert search and treeValues", func(t *testing.T) {
		array := []int{4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0, 7}
		sortArray := []int{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9}

		binaryTreeRoot := BuildTree(array)

		var arr []int
		binaryTreeRoot.TreeValues(&arr)
		assert.Equal(t, arr, sortArray)


		assert.Equal(t, binaryTreeRoot.Search(9), true)
		assert.Equal(t, binaryTreeRoot.Search(0), true)
		assert.Equal(t, binaryTreeRoot.Search(2), true)
		assert.Equal(t, binaryTreeRoot.Search(4), true)
		assert.Equal(t, binaryTreeRoot.Search(6), true)
		assert.Equal(t, binaryTreeRoot.Search(10), false)
	})

	t.Run("test insert/remove", func(t *testing.T) {
		array := []int{4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0, 7}
		sortArray := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}

		binaryTreeRoot := BuildTree(array)

		Remove(5, binaryTreeRoot)
		Remove(2, binaryTreeRoot)
		Remove(6, binaryTreeRoot)

		var arr []int
		binaryTreeRoot.TreeValues(&arr)

		assert.Equal(t, arr, sortArray)


		binaryTreeRoot.Insert(5)
		binaryTreeRoot.Insert(6)
		binaryTreeRoot.Insert(2)

		RemoveWithoutReqursion(5, binaryTreeRoot)
		RemoveWithoutReqursion(2, binaryTreeRoot)
		RemoveWithoutReqursion(6, binaryTreeRoot)

		var arr2 []int
		binaryTreeRoot.TreeValues(&arr2)

		assert.Equal(t, arr2, sortArray)
	})
}
