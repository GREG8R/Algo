package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunTest(path string, f func(string) []uint64) {
	for i := 0; i <= 9; i++ {
		inputPath := fmt.Sprintf("%stest.%d.in", path, i)
		outputPath := fmt.Sprintf("%stest.%d.out", path, i)
		if run(inputPath, outputPath, f) {
			fmt.Printf("test %d PASS\n", i)
		}
	}
}

func run(inputPath, outputPath string, f func(string) []uint64) bool {
	input, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := input.Close(); err != nil {
			panic(err)
		}
	}()

	output, err := os.Open(outputPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	scannerIn := bufio.NewScanner(input)
	scannerIn.Scan()
	inputString := scannerIn.Text()

	result := f(inputString)

	scannerOut := bufio.NewScanner(output)
	i := 0
	for scannerOut.Scan() {
		value := strconv.FormatUint(result[i], 10)
		text := scannerOut.Text()
		if text != value {
			fmt.Printf("fail expect = %s, got = %s\n", text, value)
			return false
		}
		i++
	}
	return true
}
