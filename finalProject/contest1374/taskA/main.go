package main

import "fmt"

func main(){
	var x, y, n, count int
	_, _ = fmt.Scanln(&count)
	for i := 0; i < count; i++{
		_, _ = fmt.Scanln(&x, &y, &n)

		k := n - n % x + y
		if k <= n {
			fmt.Println(k)
		} else {
			fmt.Println(k - x)
		}
	}
}
