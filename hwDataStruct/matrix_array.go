package main

type MatrixArray struct{
	length int
	lengthPart int
	countOfPart int
	array [][]interface{}
	factorResize int
}

func (a *MatrixArray) Add(element interface{}, index int) {
	if a.length == a.countOfPart * a.lengthPart {
		newArray := make([][]interface{}, a.countOfPart * a.factorResize)
		for i := 0; i < a.countOfPart * a.factorResize; i++ {
			newArray[i] = make([]interface{}, a.lengthPart)
		}
		for i := a.length; i > index; i--{
			SetElement(newArray, i, GetElement(a.array, i - 1, a.lengthPart), a.lengthPart)
		}
		SetElement(newArray, index, element, a.lengthPart)
		for i := index - 1; i >= 0; i--{
			SetElement(newArray, i, GetElement(a.array, i, a.lengthPart), a.lengthPart)
		}
		a.array = newArray
		a.length++
		a.countOfPart = (a.length / a.lengthPart) + 1
	} else {
		for i := a.length; i > index; i--{
			SetElement(a.array, i, GetElement(a.array, i - 1, a.lengthPart), a.lengthPart)
		}
		SetElement(a.array, index, element, a.lengthPart)
		a.length++
	}
}

func (a *MatrixArray) Remove(index int) interface{}{
	var removeElement = GetElement(a.array, index, a.lengthPart)
	a.length--
	for i := index; i < a.length; i++ {
		SetElement(a.array, i, GetElement(a.array, i + 1, a.lengthPart), a.lengthPart)
	}
	SetElement(a.array, a.length, nil, a.lengthPart)
	return removeElement
}

func SetElement(a [][]interface{}, index int, element interface{}, lengthPart int){
	part := index / lengthPart
	ind := index % lengthPart
	a[part][ind] = element
}

func GetElement(a [][]interface{}, index int, lengthPart int) interface{}{
	part := index / lengthPart
	ind := index % lengthPart
	return a[part][ind]
}

func (a *MatrixArray) GetLength() int{
	return a.length
}

func (a *MatrixArray) GetArray() []interface{}{
	newArray := make([]interface{}, a.length)
	for i := 0; i < a.length; i ++{
		newArray[i] = GetElement(a.array, i, a.lengthPart)
	}
	return newArray
}
