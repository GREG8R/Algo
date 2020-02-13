package main

import "os"

func CountingSortFile (in, out *os.File, n int){
	countArr := make([]int, 65536)

	for i := 0; i < n; i++{
		currentNum := ReadNum(in, i)
		countArr[currentNum]++
	}

	k := 0
	for index, num := range countArr{
		for j := 0; j < num; j++{
			WriteNum(out, k, uint16(index))
			k++
		}
	}
}

func ReadNum(f *os.File, k int) uint16{
	res := make([]byte, 2)
	offset := int64(k * 2)
	f.ReadAt(res, offset)
	result := uint16(res[0]) * uint16(256) + uint16(res[1])
	return result
}

func WriteNum(f *os.File, index int, num uint16) {
	offset := int64(index * 2)
	f.WriteAt([]byte{ byte(num / 256), byte(num % 256)}, offset)
}

func TimSort(in, out *os.File, left, right int){
	if left >= right { return }
	center := left + (right - left) / 2
	if right - left < 1024 {
		arr := make([]uint16, right - left + 1)
		for i := left; i <= right; i++{
			r := ReadNum(in, i)
			arr[i - left] = r
		}

		Qsort(arr, 0, len(arr) - 1)

		for j := left; j <= right; j++{
			current := arr[j - left]
			WriteNum(in, j, current)
		}
	} else {
		TimSort(in, out, left, center)
		TimSort(in, out, center + 1, right)
		mergeFile(in, out, left, center, right)
	}
}

func mergeFile(inputFile, resFile *os.File, left, center, right int){
	a := left
	b := center + 1
	r := 0
	for a <= center && b <= right {
		numA := ReadNum(inputFile, a)
		numB := ReadNum(inputFile, b)
		if numA < numB{
			WriteNum(resFile, r, numA)
			r++
			a++
		} else {
			WriteNum(resFile, r, numB)
			r++
			b++
		}
	}
	for a <= center {
		numA := ReadNum(inputFile, a)
		WriteNum(resFile, r, numA)
		r++
		a++
	}
	for b <= right {
		numB := ReadNum(inputFile, b)
		WriteNum(resFile, r, numB)
		r++
		b++
	}

	for j := left; j <= right; j++{
		current := ReadNum(resFile, j - left)
		WriteNum(inputFile, j, current)
	}
}

func Qsort(arr []uint16, left, right int){
	if left >= right { return }
	center := partition(arr, left, right)
	Qsort(arr, left, center - 1)
	Qsort(arr, center + 1, right)
}

func partition(arr []uint16, left, right int) int{
	i := left - 1
	pivot := arr[right]
	for j := left; j <= right; j++{
		if arr[j] <= pivot{
			i++
			x := arr[i]
			arr[i] = arr[j]
			arr[j] = x
		}
	}
	return i
}

func CountingSort (arr []uint16) []uint16{
	countArr := make([]int, 65536)
	result := make([]uint16, len(arr))

	for i := 0; i < len(arr); i++{
		countArr[arr[i]]++
	}

	k := 0
	for index, num := range countArr{
		for j := 0; j < num; j++{
			result[k] = uint16(index)
			k++
		}
	}
	return result
}
