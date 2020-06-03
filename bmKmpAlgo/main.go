package main

import "fmt"

func main(){
	text := "SOMESTRING"
	pattern := "STRING"

	fmt.Println(bm(text, pattern))
	fmt.Println(kmp(text, pattern))

	text = "SOMErrSOMESTRINGSOME"
	pattern = "SOMESTRING"

	fmt.Println(bm(text, pattern))
	fmt.Println(kmp(text, pattern))
}
