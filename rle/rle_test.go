package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_Run(t *testing.T) {
	t.Run("test RLE", func(t *testing.T) {

		inputArray1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		inputArray2 := []int{1, 2, 3, 1, 2, 3}
		inputArray3 := []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 1, 1, 1, 1, 1}

		outputArray1 := []int{11, 1}
		outputArray2 := []int{1, 1, 1, 2, 1, 3, 1, 1, 1, 2, 1, 3}
		outputArray3 := []int{4, 1, 4, 2, 5, 3, 5, 1}

		result1 := rle(inputArray1)
		result2 := rle(inputArray2)
		result3 := rle(inputArray3)

		for i, v := range outputArray1{
			assert.Equal(t, v, result1[i])
		}

		for i, v := range outputArray2{
			assert.Equal(t, v, result2[i])
		}

		for i, v := range outputArray3{
			assert.Equal(t, v, result3[i])
		}
	})
}