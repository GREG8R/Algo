package main

type FactorArray struct {
	length       int
	array        []interface{}
	factorResize int
}

func (a *FactorArray) Add(element interface{}, index int) {
	if len(a.array) == a.length {
		length := a.length + 1
		newArray := make([]interface{}, a.length*a.factorResize)
		if a.length == 0 {
			newArray = make([]interface{}, 1)
		}
		for i := 0; i < index; i++ {
			newArray[i] = a.array[i]
		}
		newArray[index] = element
		for i := index + 1; i < length; i++ {
			newArray[i] = a.array[i-1]
		}
		a.array = newArray
		a.length = length
	} else {
		for i := a.length; i > index; i-- {
			a.array[i] = a.array[i-1]
		}
		a.array[index] = element
		a.length++
	}
}

func (a *FactorArray) Remove(index int) interface{} {
	var removeElement = a.array[index]
	a.length--
	for i := 0; i < a.length; i++ {
		if i >= index {
			a.array[i] = a.array[i+1]
		}
	}
	a.array[a.length] = nil
	return removeElement
}

func (a *FactorArray) GetLength() int {
	return a.length
}

func (a *FactorArray) GetArray() []interface{} {
	return a.array
}
