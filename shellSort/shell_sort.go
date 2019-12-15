package main

import (
	"math"
)

func ShellSort(array []int, gaps []int){
	n := len(array)
	for _, gap := range gaps {
		for i := 0; i+ gap < n; i++{
			j := i + gap
			tmp := array[j]
			for ; j - gap >= 0 && array[j - gap] > tmp; j -= gap{
				array[j] = array[j - gap]
			}
			array[j] = tmp
		}
	}
}

// experimentally
func CiuraGap(n int) []int{
	gaps := []int{1, 4, 10, 23, 57, 132, 301, 701}
	reverseResult := make([]int, 0, 8)
	for _, gap := range gaps{
		if gap > n / 2 {
			break
		}
		reverseResult = append(reverseResult, gap)
	}

	length := len(gaps)
	result := make([]int, length)
	for i := length - 1; i >= 0; i--{
		result[length - i - 1] = gaps[i]
	}
	return result
}

// Sedgewick 1982 4^k+3*2^(k-1)+1
func SedgewickGap(n int) []int{
	gaps := make([]int, 0, 10)
	gap := 1
	for i := 1; gap <= n / 2; i++{
		gaps = append(gaps, gap)
		gap = int(math.Pow(float64(4), float64(i)) + float64(3) * math.Pow(float64(2), float64(i - 1)) + float64(1))
	}
	length := len(gaps)
	result := make([]int, length)
	for i := length - 1; i >= 0; i--{
		result[length - i - 1] = gaps[i]
	}
	return result
}

// shell
func SimpleGap(n int) []int{
	gaps := make([]int, 0, 10)
	gap := n / 2
	for gap > 0 {
		gaps = append(gaps, gap)
		gap /= 2
	}
	return gaps
}