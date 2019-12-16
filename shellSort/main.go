package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	runTest(2000)
	runTest(20000)
	runTest(2000000)
	runTest(20000000)
}

func runTest(n int) {
	// prepare data
	arrayBeforeSort1 := generateArray(n)
	sortArray := make([]int, n)
	mixTenPercentArray1 := make([]int, n)
	mixFiveElementsArray1 := make([]int, n)
	copy(sortArray, arrayBeforeSort1)
	simpleGaps := SimpleGap(n)
	sedgewickGaps := SedgewickGap(n)
	ciuraGaps := CiuraGap(n)

	ShellSort(sortArray, simpleGaps)
	copy(mixTenPercentArray1, sortArray)
	copy(mixFiveElementsArray1, sortArray)
	mixArray(mixTenPercentArray1, n, n/10)
	mixArray(mixFiveElementsArray1, n, 5)

	arrayBeforeSort2 := make([]int, n)
	mixTenPercentArray2 := make([]int, n)
	mixFiveElementsArray2 := make([]int, n)
	copy(arrayBeforeSort2, arrayBeforeSort1)
	copy(mixTenPercentArray2, mixTenPercentArray1)
	copy(mixFiveElementsArray2, mixFiveElementsArray1)

	arrayBeforeSort3 := make([]int, n)
	mixTenPercentArray3 := make([]int, n)
	mixFiveElementsArray3 := make([]int, n)
	copy(arrayBeforeSort3, arrayBeforeSort1)
	copy(mixTenPercentArray3, mixTenPercentArray1)
	copy(mixFiveElementsArray3, mixFiveElementsArray1)

	fmt.Printf("test for n = %d\n", n)
	// simple gaps
	fmt.Println("simple gaps shell sort")
	t := time.Now()
	ShellSort(arrayBeforeSort1, simpleGaps)
	fmt.Printf("random array, time: %d\n", getTime(n, t))

	t = time.Now()
	ShellSort(mixTenPercentArray1, simpleGaps)
	fmt.Printf("mix ten percent array, time: %d\n", getTime(n, t))

	t = time.Now()
	ShellSort(mixFiveElementsArray1, simpleGaps)
	fmt.Printf("mix five elements array, time: %d\n\n", getTime(n, t))

	// sedgewick gaps
	fmt.Println("sedgewick gaps shell sort")
	t = time.Now()
	ShellSort(arrayBeforeSort2, sedgewickGaps)
	fmt.Printf("random array, time: %d\n", getTime(n, t))

	t = time.Now()
	ShellSort(mixTenPercentArray2, sedgewickGaps)
	fmt.Printf("mix ten percent array, time: %d\n", getTime(n, t))

	t = time.Now()
	ShellSort(mixFiveElementsArray2, sedgewickGaps)
	fmt.Printf("mix five elements array, time: %d\n\n", getTime(n, t))

	if n > 2000000 {
		return
	}

	// ciura gaps
	fmt.Println("ciura gaps shell sort")
	t = time.Now()
	ShellSort(arrayBeforeSort3, ciuraGaps)
	fmt.Printf("random array, time: %d\n", getTime(n, t))

	t = time.Now()
	ShellSort(mixTenPercentArray3, ciuraGaps)
	fmt.Printf("mix ten percent array, time: %d\n", getTime(n, t))

	t = time.Now()
	ShellSort(mixFiveElementsArray3, ciuraGaps)
	fmt.Printf("mix five elements array, time: %d\n\n", getTime(n, t))
}

func getTime(n int, t time.Time) int64 {
	if n <= 1000 {
		return time.Now().Sub(t).Nanoseconds()
	}
	if n <= 20000 {
		return time.Now().Sub(t).Microseconds()
	}
	return time.Now().Sub(t).Milliseconds()
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
