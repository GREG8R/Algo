package main

type VectorArray struct{
	length int
	array []interface{}
	capResize int
}

func (a *VectorArray) Add(element interface{}, index int){
	if len(a.array) == a.length {
		length := a.length + 1
		newArray := make([]interface{}, a.length + a.capResize)
		for i := 0; i < length; i++{
			if i < index {
				newArray[i] = a.array[i]
			}
			if i == index{
				newArray[i] = element
			}
			if i > index {
				newArray[i] = a.array[i - 1]
			}
		}
		a.array = newArray
		a.length = length
	} else {
		for i := a.length; i > index; i--{
			a.array[i] = a.array[i - 1]
		}
		a.array[index] = element
		a.length++
	}
}

func (a *VectorArray) Remove(index int) interface{}{
	var removeElement = a.array[index]
	a.length--
	for i := 0; i < a.length; i++ {
		if i >= index {
			a.array[i] = a.array[i + 1]
		}
	}
	a.array[a.length] = nil
	return removeElement
}

func (a *VectorArray) GetLength() int{
	return a.length
}

func (a *VectorArray) GetArray() []interface{}{
	return a.array
}
