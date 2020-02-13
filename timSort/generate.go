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

func GenerateFileForTest() ([]uint16, *os.File, *os.File){
	result := make([]uint16, 10000)
	out, err := os.Create("testFile")
	if err != nil {
		panic(err)
	}

	res, err := os.Create("resFile")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10000; i++{
		result[i] = uint16(rand.Intn(65535))
		out.Write([]byte{ byte(result[i] / 256), byte(result[i] % 256)})
	}
	return result, out, res
}

func CloseFiles(in, out *os.File){
	in.Close()
	out.Close()
}
