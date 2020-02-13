package main

import "os"

func MergeSortFile(input string, left, right int){
	if left >= right { return }
	center := left + (right - left) / 2

	//if right - left < 1024 {
	//
	//} else {
		MergeSortFile(input, left, center)
		MergeSortFile(input, center + 1, right)
		MergeFile(input, left, center, right)
	//}
}

func MergeFile(input string, left, center, right int){
	res := "resFile"
	resFile, err := os.Create(res)
	if err != nil {
		panic(err)
	}
	defer resFile.Close()

	inputFile, err := os.OpenFile(input, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

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

//func WriteNum(f *os.File, num uint16) {
//	f.Write([]byte{ byte(num / 256), byte(num % 256)})
//}

//func ReadNum(fileName string, k int) uint16{
//	readFile, err := os.Open(fileName)
//	if err != nil {
//		panic(err)
//	}
//	defer readFile.Close()
//	res := make([]byte, 2)
//	offset := int64(k * 2)
//	readFile.ReadAt(res, offset)
//	result := uint16(res[0]) * uint16(256) + uint16(res[1])
//	return result
//}

//func WriteNum(fileName string, num uint16) {
//	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	f.Write([]byte{ byte(num / 256), byte(num % 256)})
//}

//func WriteNumOfIndex(fileName string, index int, num uint16) {
//	f, err := os.OpenFile(fileName, os.O_WRONLY, 0600)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	offset := int64(index * 2)
//	f.WriteAt([]byte{ byte(num / 256), byte(num % 256)}, offset)
//}




func MergeSort(arr []uint16, left, right int){
	if left >= right { return }
	center := left + (right - left) / 2

	MergeSort(arr, left, center)
	MergeSort(arr, center + 1, right)
	Merge(arr, left, center, right)
}

func Merge(arr []uint16, left, center, right int){
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
