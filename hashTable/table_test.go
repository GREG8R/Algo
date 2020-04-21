package main

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_Run(t *testing.T) {

	t.Run("test hash table", func(t *testing.T) {
		array := []int{4, 64, 54, 9, 19, 49, 69, 8, 0}

		ht := Init(10)

		for _, i := range array{
			ht.Set(i, fmt.Sprintf("some_value_%d", i))
		}

		ht.Delete(9)
		assert.Equal(t, ht.nodes[9].key, 9)
		assert.Equal(t, ht.nodes[9].isDeleted, true)

		for _, i := range array {
			assert.Equal(t, ht.Get(i), fmt.Sprintf("some_value_%d", i))
		}

		array2 := []int{44, 464, 454, 49, 419, 429, 469, 48, 40}

		for _, i := range array2{
			ht.Set(i, fmt.Sprintf("some_value_%d", i))
		}

		ht.Set(9, "some_value_9")
		for _, i := range array {
			assert.Equal(t, ht.Get(i), fmt.Sprintf("some_value_%d", i))
		}

		for _, i := range array2 {
			assert.Equal(t, ht.Get(i), fmt.Sprintf("some_value_%d", i))
		}
	})
}
