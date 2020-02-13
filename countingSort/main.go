package main

import (
	"fmt"
	"time"
)

func main(){
	n := 1000000
	GoCounting(n)
	GoTim(n)
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

func GoCounting(n int){
	in, out := GenerateFile(n)
	fmt.Printf("test for n = %d\n", n)
	fmt.Println("counting sort")
	t := time.Now()
	CountingSortFile(in, out, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())
	CloseFiles(in, out)
}
