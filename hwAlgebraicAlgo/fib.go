package main

import "math"

func FibRecursive(n int) uint64 {
	if n <= 2 {
		return 1
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

func FibIteration(n int) uint64 {
	if n <= 2 {
		return 1
	}

	prev := uint64(1)
	result := uint64(1)
	for i := 2; i < n; i++ {
		buf := result
		result = prev + result
		prev = buf
	}
	return result
}

func FibFormula(n int) uint64 {
	if n <= 2 {
		return 1
	}

	fi := (1 + math.Sqrt(float64(5))) / 2
	result := math.Pow(fi, float64(n))/math.Sqrt(float64(5)) + 1/2
	return uint64(math.Round(result))
}

func FibMatrix(n int) uint64 {
	if n <= 2 {
		return 1
	}

	res := []uint64{1, 0, 1, 0}
	base := []uint64{1, 1, 1, 0}
	for n > 1 {
		if n%2 == 1 {
			res = composition(res, base)
		}
		base = composition(base, base)
		n /= 2
	}
	if n > 0 {
		res = composition(res, base)
	}
	return res[1]
}

func composition(a, b []uint64) []uint64 {
	result := make([]uint64, 4)

	result[0] = a[0]*b[0] + a[1]*b[2]
	result[1] = a[0]*b[1] + a[1]*b[3]
	result[2] = a[2]*b[0] + a[3]*b[2]
	result[3] = a[2]*b[1] + a[3]*b[3]

	return result
}
