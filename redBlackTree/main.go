package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	n := 10000

	randomTree := GenerateRandomRB(n)
	randomSortTree := GenerateSortRB(n)

	fmt.Printf("test for n = %d\n", n)
	fmt.Println("search for random RBtree")
	t := time.Now()
	SearchInRBTree(randomTree.Root, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("search for random sort RBtree")
	t = time.Now()
	SearchInRBTree(randomSortTree.Root, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	randomTreap := GenerateRandom(n)
	randomSortTreap := GenerateSort(n)

	fmt.Println("search for random treap")
	t = time.Now()
	SearchInTreap(randomTreap, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

	fmt.Println("search for random sort treap")
	t = time.Now()
	SearchInTreap(randomSortTreap, n)
	fmt.Printf("time: %d\n", time.Now().Sub(t).Microseconds())

}

func SearchInRBTree(root *RedBlackNode, n int) {
	for i := 0; i < n/10; i++ {
		r := rand.Intn(n)
		root.Search(r)
	}
}

func SearchInTreap(root *Treap, n int) {
	for i := 0; i < n/10; i++ {
		r := rand.Intn(n)
		root.Search(r)
	}
}
