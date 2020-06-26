package main

func rle(arr []int) []int{
	result := make([]int, 0, 0)
	lastVal := arr[0]
	lastIndex := 0
	result = append(result, 1)
	result = append(result, lastVal)

	for i := 1; i < len(arr); i++ {
		if arr[i] == lastVal {
			result[lastIndex]++
		} else {
			lastIndex += 2
			lastVal = arr[i]
			result = append(result, 1)
			result = append(result, lastVal)
		}
	}
	return result
}

