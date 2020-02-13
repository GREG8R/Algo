package main

import (
	"math/rand"
	"os"
)

func GenerateFile(n int) {
	out, err := os.Create("inputFile")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	for k := 0; k < n; k++{
		k := uint16(rand.Intn(65535))
		result := []byte{ byte(k / 256), byte(k % 256)}
		out.Write(result)
	}
}

func GenerateFileForTest() []uint16{
	result := make([]uint16, 100)
	out, err := os.Create("testFile")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	for i := 0; i < 100; i++{
		result[i] = uint16(rand.Intn(65535))
		out.Write([]byte{ byte(result[i] / 256), byte(result[i] % 256)})
	}
	return result
}
