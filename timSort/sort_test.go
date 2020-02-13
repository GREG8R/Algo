package main

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestInputOutput_Run(t *testing.T) {
	t.Run("test read write file", func(t *testing.T){
		res, err := os.Open("testFile")
		if err != nil {
			panic(err)
		}
		inputArray := GenerateFileForTest()
		for i, k := range inputArray{
			r := ReadNum(res, i)
			fmt.Println(r)
			assert.Equal(t, k, r)
		}

		res.Close()
	})
}

func TestMergeSort_Run(t *testing.T) {
	t.Run("sort test", func(t *testing.T) {
		array := []uint16{7, 4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0}
		MergeSort(array, 0, len(array) - 1)
		assert.Equal(t, array, []uint16{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9})
	})

	t.Run("test write in file sort", func(t *testing.T){
		inputArray := GenerateFileForTest()
		input := "testFile"

		MergeSortFile(input, 0, len(inputArray) - 1)
		MergeSort(inputArray, 0, len(inputArray) - 1)

		inputFile, err := os.Open(input)
		if err != nil {
			panic(err)
		}

		for i, k := range inputArray{
			r := ReadNum(inputFile, i)
			fmt.Println(r)
			assert.Equal(t, k, r)
		}

		inputFile.Close()
	})
}