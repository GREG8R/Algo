package main

//enqueue(int priority, T item) - поместить элемент в очередь
//T dequeue() - выбрать элемент из очереди

type Queue struct{
	array []*QueueWithPriority
	last int // индекс последнего заполненного элемента
	indexOfPriorityForGet int // индекс в массиве для очереди с самым высоким приоритетом
	IndexOfMaxPriority int // индекс самого низкого приоритета
}

func (q *Queue) Enqueue(priority int, element interface{}){
	for i, value := range q.array {
		if value.Priority == priority{
			q.array[i].Push(element)
			if priority < q.array[q.indexOfPriorityForGet].Priority{
				q.indexOfPriorityForGet = i
			}
			if priority > q.array[q.IndexOfMaxPriority].Priority{
				q.IndexOfMaxPriority = i
			}
			return
		}
	}

	if q.last + 1 == len(q.array){
		newArray := make([]*QueueWithPriority, len(q.array) * 2)
		if len(q.array) == 0{
			newArray = make([]*QueueWithPriority, 1)
		}
		for i := 0; i < len(q.array); i++ {
			newArray[i] = q.array[i]
		}
		q.array = newArray
	}
	q.last++
	q.array[q.last] = &QueueWithPriority{
		Priority:priority,
	}
	q.array[q.last].array = []interface{}{}
	q.array[q.last].Push(element)
	if priority < q.array[q.indexOfPriorityForGet].Priority{
		q.indexOfPriorityForGet = q.last
	}
	if priority > q.array[q.IndexOfMaxPriority].Priority{
		q.IndexOfMaxPriority = q.last
	}
}

func (q *Queue) Dequeue() interface{}{
	element := q.array[q.indexOfPriorityForGet].Pop()
	if element != nil {
		return element
	}
	q.array[q.indexOfPriorityForGet].IsEmpty = true

	q.indexOfPriorityForGet = q.IndexOfMaxPriority
	for i := 0; i <= q.last; i++{
		if !q.array[i].IsEmpty && q.array[q.indexOfPriorityForGet].Priority > q.array[i].Priority{
			q.indexOfPriorityForGet = i
		}
	}

	return q.Dequeue()
}

type QueueWithPriority struct{
	BufferedArray
	Priority int
	IsEmpty bool
}

// вспомогательная структура зацикленный массив, для очереди, добавлять в конец вытаскивать с начала.
type BufferedArray struct{
	array []interface{}
	first int
	last int
	length int
}

const(
	resizeValue = 2
)

func (a *BufferedArray) Push(element interface{}){
	length := a.length
	if length != 0 && a.first != (a.last + 1) % length{
		a.array[a.last] = element
		a.last = (a.last + 1) % length
		return
	}

	newArray := make([]interface{}, length * resizeValue)
	if length == 0{
		newArray = make([]interface{}, 2)
	}
	for i := a.first; i < length; i++{
		newArray[i - a.first] = a.array[i]
	}
	if a.last < a.first{
		for i := 0; i < a.first; i++{
			newArray[length - a.first] = a.array[i]
		}
		a.first = 0
		a.last = length - 1
	}
	a.array = newArray
	a.length = len(newArray)
	a.array[a.last] = element
	a.last = (a.last + 1) % a.length
}

func (a *BufferedArray) Pop() interface{}{
	if a.first == a.last {
		return nil
	}

	element := a.array[a.first]
	a.array[a.first] = nil
	a.first = (a.first + 1) % a.length
	return element
}
