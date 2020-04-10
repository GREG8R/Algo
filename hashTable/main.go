package main

import (
	"fmt"
)

func main() {
	ht := Init(100)
	array := []int{4, 64, 54, 9, 19, 29, 69, 8, 0}

	for _, i := range array{
		ht.Set(i, fmt.Sprintf("some_value_%d", i))
	}

	fmt.Println("HashTable print")
	for _, i := range array {
		fmt.Printf("key = %d, value = %s\n", i, ht.Get(i))
	}
}
