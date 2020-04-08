package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const ACommand string = "A"
const QCommand string = "Q"

func main() {
	input, err := os.Open("sum.in")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := input.Close(); err != nil {
			panic(err)
		}
	}()

	output, err := os.Create("sum.out")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	firstString := scanner.Text()

	splitString := strings.Split(firstString, " ")

	countOfNumbers, err := strconv.Atoi(splitString[0])
	if err != nil {
		panic("NAN")
	}

	tree := buildTreeOfSegments(countOfNumbers)

	countOfCommand, err := strconv.Atoi(splitString[1])
	if err != nil {
		panic("NAN")
	}

	for i := 0; i < countOfCommand; i++ {
		scanner.Scan()
		splitStr := strings.Split(scanner.Text(), " ")

		if splitStr[0] == ACommand {
			index, err := strconv.Atoi(splitStr[1])
			if err != nil {
				panic("NAN")
			}
			value, err := strconv.Atoi(splitStr[2])
			if err != nil {
				panic("NAN")
			}

			aFunc(index, uint64(value), tree)
		}

		if splitStr[0] == QCommand {
			left, err := strconv.Atoi(splitStr[1])
			if err != nil {
				panic("NAN")
			}
			right, err := strconv.Atoi(splitStr[2])
			if err != nil {
				panic("NAN")
			}

			sum := qFunc(left, right, tree)
			_, err = output.WriteString(strconv.FormatUint(sum, 10))
			if err != nil {
				panic("NAN")
			}
			_, err = output.WriteString("\n")
			if err != nil {
				panic("NAN")
			}
		}
	}
}

func buildTreeOfSegments(count int) []uint64 {
	i := 1
	for i < count {
		i *= 2
	}
	return make([]uint64, i*2)
}

func aFunc(index int, value uint64, tree []uint64) {
	indexOfChange := len(tree)/2 + index - 1
	tree[indexOfChange] = value

	for indexOfChange > 1 {
		parent := indexOfChange / 2
		tree[parent] = tree[parent*2] + tree[parent*2+1]
		indexOfChange = parent
	}
}

func qFunc(left, right int, tree []uint64) uint64 {
	sum := uint64(0)

	left = len(tree)/2 + left - 1
	right = len(tree)/2 + right - 1

	for left <= right {
		if left%2 != 0 {
			sum += tree[left]
		}
		if right%2 == 0 {
			sum += tree[right]
		}

		left = (left + 1) / 2
		right = (right - 1) / 2
	}

	return sum
}
