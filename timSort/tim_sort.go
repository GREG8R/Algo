package main

import "os"

func MergeSortFile(in, out *os.File, left, right int){
	if left >= right { return }
	center := left + (right - left) / 2
	MergeSortFile(in, out, left, center)
	MergeSortFile(in, out, center + 1, right)
	mergeFile(in, out, left, center, right)
}

func MergeSortFileWithOptimisation(in, out *os.File, left, right int){
	if left >= right { return }
	center := left + (right - left) / 2
	if right - left < 1024 {
		arr := make([]uint16, right - left + 1)
		for i := left; i <= right; i++{
			r := ReadNum(in, i)
			arr[i - left] = r
		}
		MergeSort(arr, 0, len(arr) - 1)
		for j := left; j <= right; j++{
			current := arr[j - left]
			WriteNum(in, j, current)
		}
	} else {
		MergeSortFileWithOptimisation(in, out, left, center)
		MergeSortFileWithOptimisation(in, out, center + 1, right)
		mergeFile(in, out, left, center, right)
	}
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

func TimHeapSort(in, out *os.File, left, right int){
	if left >= right { return }
	center := left + (right - left) / 2
	if right - left < 1024 {
		arr := make([]uint16, right - left + 1)
		for i := left; i <= right; i++{
			r := ReadNum(in, i)
			arr[i - left] = r
		}

		HeapSort(arr)

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

func HeapSort(array []uint16) {
	swap := func(x, y int) {
		tmp := array[x]
		array[x] = array[y]
		array[y] = tmp
	}

	down := func(size, root int) {}
	down = func(size, root int) {
		l := root*2 + 1
		r := l + 1
		x := root
		if l < size && array[l] > array[x] {
			x = l
		}
		if r < size && array[r] > array[x] {
			x = r
		}
		if x == root {
			return
		}
		swap(root, x)
		down(size, x)
	}

	for node := len(array)/2 - 1; node >= 0; node-- {
		down(len(array), node)
	}

	down(len(array), 0)
	for i := len(array) - 1; i >= 0; i-- {
		swap(i, 0)
		down(i, 0)
	}
}

func MergeSort(arr []uint16, left, right int){
	if left >= right { return }
	center := left + (right - left) / 2

	MergeSort(arr, left, center)
	MergeSort(arr, center + 1, right)
	merge(arr, left, center, right)
}

func merge(arr []uint16, left, center, right int){
	res := make([]uint16, right - left + 1)
	a := left
	b := center + 1
	r := 0
	for a <= center && b <= right {
		if arr[a] < arr[b]{
			res[r] = arr[a]
			r++
			a++
		} else {
			res[r] = arr[b]
			r++
			b++
		}
	}
	for a <= center {
		res[r] = arr[a]
		r++
		a++
	}
	for b <= right {
		res[r] = arr[b]
		r++
		b++
	}

	for j := left; j <= right; j++{
		arr[j] = res[j - left]
	}
}
