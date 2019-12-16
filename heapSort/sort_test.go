package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShellSort_Run(t *testing.T) {
	t.Run("sort test", func(t *testing.T) {
		array := []int{7, 4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0}
		HeapSort(array)
		assert.Equal(t, array, []int{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9})
	})
}
