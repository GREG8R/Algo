package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInputOutput_Run(t *testing.T) {
	t.Run("test read write file", func(t *testing.T){
		inputArray, in, out := GenerateFileForTest()
		for i, k := range inputArray{
			r := ReadNum(in, i)
			assert.Equal(t, k, r)
		}

		CloseFiles(in, out)
	})
}

func TestMergeSort_Run(t *testing.T) {
	t.Run("sort test", func(t *testing.T) {
		array := []uint16{7, 4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0}
		MergeSort(array, 0, len(array) - 1)
		assert.Equal(t, array, []uint16{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9})
	})

	t.Run("test MergeSortFile", func(t *testing.T){
		inputArray, in, out := GenerateFileForTest()

		MergeSortFile(in, out, 0, len(inputArray) - 1)
		MergeSort(inputArray, 0, len(inputArray) - 1)

		for i, k := range inputArray{
			r := ReadNum(in, i)
			assert.Equal(t, k, r)
		}

		CloseFiles(in, out)
	})

	t.Run("test MergeSortFileWithOptimisation", func(t *testing.T){
		inputArray, in, out := GenerateFileForTest()

		MergeSortFileWithOptimisation(in, out, 0, len(inputArray) - 1)
		MergeSort(inputArray, 0, len(inputArray) - 1)

		for i, k := range inputArray{
			r := ReadNum(in, i)
			assert.Equal(t, k, r)
		}

		CloseFiles(in, out)
	})

	t.Run("test TimSort", func(t *testing.T){
		inputArray, in, out := GenerateFileForTest()

		TimSort(in, out, 0, len(inputArray) - 1)
		MergeSort(inputArray, 0, len(inputArray) - 1)

		for i, k := range inputArray{
			r := ReadNum(in, i)
			assert.Equal(t, k, r)
		}

		CloseFiles(in, out)
	})
}