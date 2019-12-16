package main

type SimpleArrayOld struct {
	length int
	array  []interface{}
}

func (a *SimpleArrayOld) Add(element interface{}, index int) {
	newArray := make([]interface{}, a.length+capResizeSA)
	for i := 0; i < a.length+capResizeSA; i++ {
		if i < index {
			newArray[i] = a.array[i]
		}
		if i == index {
			newArray[i] = element
		}
		if i > index {
			newArray[i] = a.array[i-1]
		}
	}
	a.length = a.length + capResizeSA
	a.array = newArray
}

func (a *SimpleArrayOld) GetLength() int {
	return a.length
}

func (a *SimpleArrayOld) GetArray() []interface{} {
	return a.array
}

func (a *SimpleArrayOld) Remove(index int) interface{} {
	var removeElement = a.array[index]
	a.length = a.length - capResizeSA
	newArray := make([]interface{}, a.length)
	for i := 0; i < a.length; i++ {
		if i >= index {
			newArray[i] = a.array[i+1]
		} else {
			newArray[i] = a.array[i]
		}
	}
	a.array = newArray
	return removeElement
}
