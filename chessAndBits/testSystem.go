package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunTest(path string, f func(string) []uint64) {
	for i := 0; i <= 9; i++{
		inputPath := fmt.Sprintf("%stest.%d.in", path, i)
		outputPath := fmt.Sprintf("%stest.%d.out", path, i)
		if run(inputPath, outputPath, f) {
			fmt.Printf("test %d PASS\n", i)
		}
	}
}

func run(inputPath, outputPath string, f func(string) []uint64) bool{
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

func King(strPosition string) []uint64 {
	result := make([]uint64, 2)
	position, _ := strconv.Atoi(strPosition)

	k := uint64(1) << position
	nA := uint64(0xFEFEFEFEFEFEFEFE)
	nH := uint64(0x7F7F7F7F7F7F7F7F)
	p := make([]uint64, 10)
	p[1] = (k & nA) >> 9
	p[2] = k >> 8
	p[3] = (k & nH) >> 7
	p[4] = (k & nA) >> 1
	p[6] = (k & nH) << 1
	p[7] = (k & nA) << 7
	p[8] = k << 8
	p[9] = (k & nH) << 9
	result[1] = p[7] | p[8] | p[9] |
				p[4] | 	      p[6] |
				p[1] | p[2] | p[3]

	for _, v := range p {
		if v > 0 {
			result[0]++
		}
	}

	return result
}

func Knight(strPosition string) []uint64 {
	result := make([]uint64, 2)
	position, _ := strconv.Atoi(strPosition)

	k := uint64(1) << position
	nA := uint64(0xFEFEFEFEFEFEFEFE)
	nAB := uint64(0xFCFCFCFCFCFCFCFC)
	nGH := uint64(0x3F3F3F3F3F3F3F3F)
	nH := uint64(0x7F7F7F7F7F7F7F7F)
	p := make([]uint64, 10)
	p[1] = (k & nA) << 15
	p[2] = (k & nH) << 17
	p[3] = (k & nGH) << 10
	p[9] = (k & nAB) << 6

	p[4] = (k & nGH) >> 6
	p[6] = (k & nH) >> 15
	p[7] = (k & nA) >> 17
	p[8] = (k & nAB) >> 10

	result[1] = p[7] | p[8] | p[9] |
				p[4] | 	      p[6] |
				p[1] | p[2] | p[3]

	for _, v := range p {
		if v > 0 {
			result[0]++
		}
	}

	return result
}
