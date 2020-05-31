package main

import (
	"math"
)

func bm(text, pattern string) int {
	prefix := CreatePrefix(pattern)
	suffix := CreateSuffix(pattern)

	last := len(pattern) - 1
	t := 0
	for t < len(text) - last{
		p := last
		for p >= 0 && text[t + p] == pattern[p]{
			p--
		}
		if p == -1{
			return t
		}
		t += int(math.Max(float64(p - prefix[text[t + p]]), float64(suffix[p + 1])))
	}

	return -1
}

func CreatePrefix(pattern string) []int{
	prefix := make([]int, 128)
	for i := 0; i < len(prefix); i++ {
		prefix[i] = -1
	}
	for i := 0; i < len(pattern) - 1; i++{
		prefix[pattern[i]] = i
	}
	return prefix
}

func CreateSuffix(pattern string) []int{
	suffix := make([]int, len(pattern) + 1)
	for i := 0; i < len(suffix); i++ {
		suffix[i] = len(pattern)
	}
	suffix[len(pattern)] = 1

	for j := len(pattern) - 1; j >= 0; j--{
		for at := j; at < len(pattern); at++{
			a := substring(pattern, 0, at)
			for i := at - 1; i >= 0; i--{
				b := substring(pattern, i, len(a))
				if a == b {
					suffix[j] = at - 1
					at = len(pattern)
					break
				}
			}
		}
	}

	return suffix
}

func substring(str string, a, b int) string{
	arr := []byte(str)
	return string(arr[a:b])
}

func kmp(){

}

