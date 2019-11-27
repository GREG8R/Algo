package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Compare(a1 Array, a2 []interface{}) bool{
	length := a1.GetLength()
	arr := a1.GetArray()
	if length != len(a2){
		return false
	}

	for i := 0; i < length; i++{
		if arr[i] != a2[i]{
			return false
		}
	}
	return true
}

func TestSimpleArray_Run(t *testing.T){
	fmt.Println("simple array")
	t.Run("add element to array", func(t *testing.T) {
		testArray := []interface{}{1, 2, 3, 4}

		simpleArray := &SimpleArray{
			length: len(testArray),
			array: testArray,
		}

		simpleArray.Add(12, 2)
		testArray = []interface{}{1, 2, 12, 3, 4}

		res := Compare(simpleArray, testArray)
		assert.Equal(t, true, res)
	})

	t.Run("remove element", func(t *testing.T) {
		testArray := []interface{}{1, 2, 12, 3, 4}

		arr := &SimpleArray{
			length: len(testArray),
			array: testArray,
		}

		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Remove(2)
		arr.Remove(2)
		arr.Remove(2)

		removeElement := arr.Remove(2)
		assert.Equal(t, 12, removeElement)
		testArray =  []interface{}{1, 2, 3, 4}

		res := Compare(arr, testArray)
		assert.Equal(t, true, res)
	})
}

func TestVectorArray_Run(t *testing.T){
	fmt.Println("vector array")
	t.Run("add element to array", func(t *testing.T) {
		testArray := []interface{}{1, 2, 3, 4}

		arr := &VectorArray{
			length: len(testArray),
			array: testArray,
			capResize: 100,
		}

		arr.Add(12, 2)
		testArray = []interface{}{1, 2, 12, 3, 4}

		res := Compare(arr, testArray)
		assert.Equal(t, true, res)
	})

	t.Run("remove element", func(t *testing.T) {
		testArray := []interface{}{1, 2, 12, 3, 4}

		arr := &VectorArray{
			length: len(testArray),
			array: testArray,
			capResize: 100,
		}

		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Remove(2)
		arr.Remove(2)
		arr.Remove(2)

		removeElement := arr.Remove(2)
		assert.Equal(t, 12, removeElement)
		testArray = []interface{}{1, 2, 3, 4}

		res := Compare(arr, testArray)
		assert.Equal(t, true, res)
	})
}

func TestFactorArray_Run(t *testing.T){
	fmt.Println("factor array")
	t.Run("add element to array", func(t *testing.T) {
		testArray := []interface{}{1, 2, 3, 4}

		arr := &FactorArray{
			length: len(testArray),
			array: testArray,
			factorResize: 2,
		}

		arr.Add(12, 2)
		testArray = []interface{}{1, 2, 12, 3, 4}

		res := Compare(arr, testArray)
		assert.Equal(t, true, res)
	})

	t.Run("remove element", func(t *testing.T) {
		testArray := []interface{}{1, 2, 12, 3, 4}

		arr := &FactorArray{
			length: len(testArray),
			array: testArray,
			factorResize: 2,
		}

		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Remove(2)
		arr.Remove(2)
		arr.Remove(2)

		removeElement := arr.Remove(2)
		assert.Equal(t, 12, removeElement)
		testArray = []interface{}{1, 2, 3, 4}

		res := Compare(arr, testArray)
		assert.Equal(t, true, res)
	})
}

func TestMatrixArray_Run(t *testing.T){
	fmt.Println("matrix array")
	t.Run("add element to array", func(t *testing.T) {
		testArray := []interface{}{1, 2, 3, 4, 5}

		marr := make([][]interface{}, 1)
		marr[0] = testArray

		arr := &MatrixArray{
			length: 5,
			array: marr,
			factorResize: 2,
			countOfPart:1,
			lengthPart:5,
		}

		arr.Add(12, 2)
		arr.Add(12, 2)
		arr.Add(12, 2)
		testArray = []interface{}{1, 2, 12, 12, 12, 3, 4, 5}

		res := Compare(arr, testArray)
		assert.Equal(t, true, res)
	})

	t.Run("remove element", func(t *testing.T) {
		testArray := []interface{}{1, 2, 3, 4, 5}

		marr := make([][]interface{}, 1)
		marr[0] = testArray

		arr := &MatrixArray{
			length: 5,
			array: marr,
			factorResize: 2,
			countOfPart:1,
			lengthPart:5,
		}

		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Add(13, 2)
		arr.Remove(2)
		arr.Remove(2)
		arr.Remove(2)

		removeElement := arr.Remove(2)
		assert.Equal(t, 3, removeElement)
		testArray = []interface{}{1, 2, 4, 5}

		res := Compare(arr, testArray)
		assert.Equal(t, true, res)
	})
}

func TestBufferedArray_Run(t *testing.T){
	fmt.Println("buffered array")
	t.Run("push and pop elements to array", func(t *testing.T) {
		testArray := []interface{}{1, 2, 3, nil}

		arr := BufferedArray{
			array:  testArray,
			first:  0,
			last:   3,
			length: len(testArray),
		}

		arr.Push(4)
		arr.Push(5)
		arr.Push(6)
		arr.Push(7)

		assert.Equal(t, arr.last, 7)
		assert.Equal(t, arr.first, 0)

		pop := arr.Pop()
		assert.Equal(t, 1, pop)
		pop = arr.Pop()
		assert.Equal(t, 2, pop)
		pop = arr.Pop()
		assert.Equal(t, 3, pop)

		assert.Equal(t, arr.last, 7)
		assert.Equal(t, arr.first, 3)

		arr.Push(8)
		arr.Push(9)
		arr.Push(10)

		assert.Equal(t, arr.last, 2)
		assert.Equal(t, arr.first, 3)

		pop = arr.Pop()
		assert.Equal(t, 4, pop)
		pop = arr.Pop()
		assert.Equal(t, 5, pop)
		pop = arr.Pop()
		assert.Equal(t, 6, pop)
		pop = arr.Pop()
		assert.Equal(t, 7, pop)
		pop = arr.Pop()
		assert.Equal(t, 8, pop)
		pop = arr.Pop()
		assert.Equal(t, 9, pop)
		pop = arr.Pop()
		assert.Equal(t, 10, pop)

		pop = arr.Pop()
		assert.Equal(t, nil, pop)

		assert.Equal(t, arr.first, arr.last)
	})

	t.Run("push and pop elements to empty array", func(t *testing.T) {
		testArray := []interface{}{}

		arr := BufferedArray{
			array:  testArray,
			first:  0,
			last:   0,
			length: 0,
		}

		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		arr.Push(4)
		arr.Push(5)
		arr.Push(6)
		arr.Push(7)

		assert.Equal(t, arr.last, 7)
		assert.Equal(t, arr.first, 0)

		pop := arr.Pop()
		assert.Equal(t, 1, pop)
		pop = arr.Pop()
		assert.Equal(t, 2, pop)
		pop = arr.Pop()
		assert.Equal(t, 3, pop)

		assert.Equal(t, arr.last, 7)
		assert.Equal(t, arr.first, 3)

		arr.Push(8)
		arr.Push(9)
		arr.Push(10)

		assert.Equal(t, arr.last, 2)
		assert.Equal(t, arr.first, 3)

		pop = arr.Pop()
		assert.Equal(t, 4, pop)
		pop = arr.Pop()
		assert.Equal(t, 5, pop)
		pop = arr.Pop()
		assert.Equal(t, 6, pop)
		pop = arr.Pop()
		assert.Equal(t, 7, pop)
		pop = arr.Pop()
		assert.Equal(t, 8, pop)
		pop = arr.Pop()
		assert.Equal(t, 9, pop)
		pop = arr.Pop()
		assert.Equal(t, 10, pop)

		pop = arr.Pop()
		assert.Equal(t, nil, pop)

		assert.Equal(t, arr.first, arr.last)
	})
}

func TestPriorityQueue_Run(t *testing.T){
	fmt.Println("priority queue")
	t.Run("enqueue and dequeue elements to queue", func(t *testing.T) {
		queue := Queue{
			array: []*QueueWithPriority{},
			last: -1,
			indexOfPriorityForGet: 0,
		}

		queue.Enqueue(0, 0)
		queue.Enqueue(0, 1)
		queue.Enqueue(0, 2)
		queue.Enqueue(0, 3)

		queue.Enqueue(1, 10)
		queue.Enqueue(1, 11)
		queue.Enqueue(1, 12)
		queue.Enqueue(1, 31)

		elem := queue.Dequeue()
		assert.Equal(t, elem, 0)
		elem = queue.Dequeue()
		assert.Equal(t, elem, 1)
		elem = queue.Dequeue()
		assert.Equal(t, elem, 2)
		elem = queue.Dequeue()
		assert.Equal(t, elem, 3)

		elem = queue.Dequeue()
		assert.Equal(t, elem, 10)
		elem = queue.Dequeue()
		assert.Equal(t, elem, 11)
		elem = queue.Dequeue()
		assert.Equal(t, elem, 12)
		elem = queue.Dequeue()
		assert.Equal(t, elem, 31)

		elem = queue.Dequeue()
		assert.Equal(t, elem, nil)
	})
}
