package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInputOutput_Run(t *testing.T) {
	t.Run("sort test", func(t *testing.T) {
		array := []uint16{7, 4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0}
		sortArray := []uint16{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9}
		res := CountingSort(array)
		for i, k := range sortArray{
			assert.Equal(t, k, res[i])
		}
	})
}

