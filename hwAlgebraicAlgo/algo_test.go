package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlgoPow(t *testing.T) {
	t.Run("pow of composition", func(t *testing.T) {
		var number float64
		number = 2
		pow := 4
		result := PowOfComposition(number, pow)
		assert.Equal(t, result, float64(16))
	})

	t.Run("fast pow", func(t *testing.T) {
		var number float64
		number = 12
		pow := 4
		result := FastPow(number, pow)
		assert.Equal(t, result, float64(20736))
	})

	t.Run("pow of pow two", func(t *testing.T) {
		var number float64
		number = 2
		pow := 4
		result := PowOfPowTwo(number, pow)
		assert.Equal(t, result, float64(16))
	})

	t.Run("pow of pow two", func(t *testing.T) {
		var number float64
		number = 12
		pow := 4
		result := PowOfPowTwo(number, pow)
		assert.Equal(t, result, float64(20736))
	})

	t.Run("pow = 0", func(t *testing.T) {
		var number float64
		number = 12
		pow := 0
		result := PowOfPowTwo(number, pow)
		assert.Equal(t, result, float64(1))

		result = PowOfComposition(number, pow)
		assert.Equal(t, result, float64(1))

		result = FastPow(number, pow)
		assert.Equal(t, result, float64(1))
	})
}

func TestAlgoFib(t *testing.T) {
	t.Run("fib recursive", func(t *testing.T) {
		result := FibRecursive(10)
		assert.Equal(t, result, 55)

		result = FibRecursive(1)
		assert.Equal(t, result, 1)

		result = FibRecursive(2)
		assert.Equal(t, result, 1)

		result = FibRecursive(13)
		assert.Equal(t, result, 233)
	})

	t.Run("fib iteration", func(t *testing.T) {
		result := FibIteration(10)
		assert.Equal(t, result, 55)

		result = FibIteration(1)
		assert.Equal(t, result, 1)

		result = FibIteration(2)
		assert.Equal(t, result, 1)

		result = FibIteration(13)
		assert.Equal(t, result, 233)
	})

	t.Run("fib iteration", func(t *testing.T) {
		result := FibIteration(10)
		assert.Equal(t, result, 55)

		result = FibIteration(1)
		assert.Equal(t, result, 1)

		result = FibIteration(2)
		assert.Equal(t, result, 1)

		result = FibIteration(13)
		assert.Equal(t, result, 233)
	})

	t.Run("fib formula", func(t *testing.T) {
		result := FibFormula(10)
		assert.Equal(t, result, 55)

		result = FibFormula(1)
		assert.Equal(t, result, 1)

		result = FibFormula(2)
		assert.Equal(t, result, 1)

		result = FibFormula(13)
		assert.Equal(t, result, 233)
	})

	t.Run("fib matrix", func(t *testing.T) {
		result := FibMatrix(10)
		assert.Equal(t, result, 55)

		result = FibMatrix(1)
		assert.Equal(t, result, 1)

		result = FibMatrix(2)
		assert.Equal(t, result, 1)

		result = FibMatrix(13)
		assert.Equal(t, result, 233)
	})

}
