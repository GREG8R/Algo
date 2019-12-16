package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShellSort_Run(t *testing.T) {
	t.Run("sort test", func(t *testing.T) {
		array1 := []int{7, 4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0}
		array2 := []int{7, 4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0}
		array3 := []int{7, 4, 6, 2, 3, 5, 9, 1, 2, 6, 8, 0}

		ShellSort(array1, SimpleGap(12))
		ShellSort(array2, SedgewickGap(12))
		ShellSort(array3, CiuraGap(12))

		assert.Equal(t, array1, []int{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9})
		assert.Equal(t, array2, []int{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9})
		assert.Equal(t, array3, []int{0, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9})
	})
}
