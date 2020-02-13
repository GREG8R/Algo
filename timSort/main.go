package main

import (
	"fmt"
	"time"
)

func main(){

	n := 100
	GenerateFile(n)
	input := "inputFile"
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("merge sort")
	t := time.Now()
	MergeSortFile(input, 0, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())

	n = 1000
	GenerateFile(n)
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("merge sort")
	t = time.Now()
	MergeSortFile("inputFile", 0, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())

	n = 10000
	GenerateFile(n)
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("merge sort")
	t = time.Now()
	MergeSortFile("inputFile", 0, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())

}