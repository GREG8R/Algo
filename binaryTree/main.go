package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	n := 10000

	randomTree := GenerateRandomTree(n)
	randomSortTree := GenerateRandomSortTree(n)

	fmt.Printf("test for n = %d\n", n)
	fmt.Println("search for random tree")
	t := time.Now()
	SearchInTree(randomTree, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("search for random sort tree")
	t = time.Now()
	SearchInTree(randomSortTree, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("deletion for random tree")
	t = time.Now()
	DeletionInTree(randomTree, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("deletion for random sort tree")
	t = time.Now()
	DeletionInTree(randomSortTree, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	randomTreap := GenerateRandom(n)
	randomSortTreap := GenerateSort(n)

	//fmt.Printf("test for treap n = %d\n", n)
	fmt.Println("search for random treap")
	t = time.Now()
	SearchInTreap(randomTreap, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("search for random sort treap")
	t = time.Now()
	SearchInTreap(randomSortTreap, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("deletion for random treap")
	t = time.Now()
	DeletionInTreap(randomTreap, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("deletion for random sort treap")
	t = time.Now()
	DeletionInTreap(randomSortTreap, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())
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

func SearchInTreap(root *Treap, n int){
	for i := 0; i < n / 10; i++{
		r := rand.Intn(n)
		root.Search(r)
	}
}

func DeletionInTreap(root *Treap, n int){
	for i := 0; i < n / 10; i++{
		r := rand.Intn(n)
		root = root.Remove(r)
	}
}
