package main

import (
	"fmt"
	"time"
)

func main(){
	n := 1000000
	Go(n)
	GoWithOpt(n)
	GoTim(n)
	GoTimHeap(n)
}

func Go(n int){
	in, out := GenerateFile(n)
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("merge sort")
	t := time.Now()
	MergeSortFile(in, out, 0, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())
	CloseFiles(in, out)
}

func GoWithOpt(n int){
	in, out := GenerateFile(n)
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("merge sort with optimization")
	t := time.Now()
	MergeSortFileWithOptimisation(in, out, 0, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())
	CloseFiles(in, out)
}

func GoTim(n int){
	in, out := GenerateFile(n)
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("tim sort with qsort")
	t := time.Now()
	TimSort(in, out, 0, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())
	CloseFiles(in, out)
}

func GoTimHeap(n int){
	in, out := GenerateFile(n)
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("tim sort with heap")
	t := time.Now()
	TimHeapSort(in, out, 0, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())
	CloseFiles(in, out)
}
