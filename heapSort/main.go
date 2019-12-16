package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 20000000
	// prepare data
	arrayBeforeSort1 := generateArray(n)
	sortArray := make([]int, n)
	mixFivePercentArray1 := make([]int, n)
	copy(sortArray, arrayBeforeSort1)

	ShellSort(sortArray)

	copy(mixFivePercentArray1, sortArray)
	mixArray(mixFivePercentArray1, n, n/20)

	// copy data
	arrayBeforeSort2 := make([]int, n)
	mixFivePercentArray2 := make([]int, n)
	copy(arrayBeforeSort2, arrayBeforeSort1)
	copy(mixFivePercentArray2, mixFivePercentArray1)

	fmt.Printf("test for n = %d\n", n)
	fmt.Println("shell sort")
	t := time.Now()
	ShellSort(arrayBeforeSort1)
	fmt.Printf("random array, time: %d\n", time.Now().Sub(t).Milliseconds())

	t = time.Now()
	ShellSort(mixFivePercentArray1)
	fmt.Printf("five percent mixed array, time: %d\n", time.Now().Sub(t).Milliseconds())

	fmt.Println("heap sort")
	t = time.Now()
	HeapSort(arrayBeforeSort2)
	fmt.Printf("random array, time: %d\n", time.Now().Sub(t).Milliseconds())

	t = time.Now()
	HeapSort(mixFivePercentArray2)
	fmt.Printf("five percent mixed array, time: %d\n", time.Now().Sub(t).Milliseconds())
}

func generateArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(n)
	}
	return array
}

func mixArray(array []int, n int, countOfMix int) {
	for i := 0; i < countOfMix; i++ {
		index1 := rand.Intn(n)
		index2 := rand.Intn(n)
		tmp := array[index1]
		array[index1] = array[index2]
		array[index2] = tmp
	}
}
