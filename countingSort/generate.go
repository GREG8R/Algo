package main

import (
	"math/rand"
	"os"
)

func GenerateFile(n int) (*os.File, *os.File){
	in, err := os.Create("inputFile")
	if err != nil {
		panic(err)
	}
	out, err := os.Create("resFile")
	if err != nil {
		panic(err)
	}
	for k := 0; k < n; k++{
		k := uint16(rand.Intn(65535))
		result := []byte{ byte(k / 256), byte(k % 256)}
		out.Write(result)
	}
	return in, out
}

func CloseFiles(in, out *os.File){
	in.Close()
	out.Close()
}
