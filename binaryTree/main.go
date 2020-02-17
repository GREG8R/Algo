package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	n := 100000

	randomTree := GenerateRandomTree(n)
	randomSortTree := GenerateRandomSortTree(n)

	fmt.Printf("test for n = %d\n", n)
	fmt.Println("search for random tree")
	t := time.Now()
	SearchInTree(randomTree, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())

	fmt.Println("search for random sort tree")
	t = time.Now()
	SearchInTree(randomSortTree, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())

	fmt.Println("deletion for random tree")
	t = time.Now()
	DeletionInTree(randomTree, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())

	fmt.Println("deletion for random sort tree")
	t = time.Now()
	DeletionInTree(randomSortTree, n)
	fmt.Printf("time: %f\n", time.Now().Sub(t).Seconds())
}


func SearchInTree(root *Node, n int){
	for i := 0; i < n / 10; i++{
		r := rand.Intn(n)
		root.Search(r)
	}
}

func DeletionInTree(root *Node, n int){
	for i := 0; i < n / 10; i++{
		r := rand.Intn(n)
		RemoveWithoutReqursion(r, root)
	}
}