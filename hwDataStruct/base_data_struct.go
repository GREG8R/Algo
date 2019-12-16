//Динамические массивы, неполный массив, очередь с приоритетом.
//Цель: Создание разных алгоритмов для реализации Динамического массива и сравнение их производительности. Создание приоритетной очереди или неполного массива.
//1 задание. Динамические массивы.
//Написать метод добавления и удаления элементов:
//void add(T item, int index);
//T remove(int index); // возвращает удаляемый элемент
//по индексу во все варианты динамических массивов:
//SingleArray, VectorArray, FactorArray, MatrixArray *.
//+3 байта
//
//2 задание. Таблица сравнения производительности.
//Сравнить время выполнения разных операций
//для разных массивов с разным порядком значений.
//* сделать обёртку над ArrayList() и тоже сравнить.
//Составить таблицу и приложить её скриншот.
//Сделать выводы и сформулировать их в нескольких предложениях.
//+3 байта
//
//3 задание. Приоритетная очередь (на выбор).
//Написать реализацию PriorityQueue - очередь с приоритетом.
//Варианты реализации - список списков, массив списков
//Методы к реализации:
//enqueue(int priority, T item) - поместить элемент в очередь
//T dequeue() - выбрать элемент из очереди
//+4 байта
//
//4 задание. Неполный массив (на выбор).
//Написать Реализацию класса SpaceArray массив массивов с неполным заполнением.
//Делать на основе одного из уже созданных массивов и/или списков.
//+4 байта дополнительно.
//
//Напишите сколько времени ушло на домашнее задание.
//
//Для реализации можно использовать только стандартные массивы [], созданные классы массивов и/или списков. Стандартные коллекции не используем!
//
//ВАЖНО! При размещении ответа укажите, на каком языке вы выполнили ДЗ. Это поможет нам ускорить его проверку.
//Критерии оценки: 1 задание. Динамические массивы. +3 байта
//2 задание. Таблица сравнение производительности. +3 байта
//3 задание (на выбор). Приоритетная очередь. +4 байта
//4 задание (на выбор). Класс SpaceArray +4 байта

package main

import (
	"fmt"
	"sync"
	"time"
)

type Array interface {
	Add(element interface{}, index int)
	Remove(index int) interface{}
	GetLength() int
	GetArray() []interface{}
}

func main() {
	testNewAndOldArrays()
	testAllArraysParallels()
	testAllArrayAddInMiddle()
}

func testAllArrayAddInMiddle() {
	fmt.Println("Test all arrays")
	simple := SimpleArray{
		length: 0,
		array:  []interface{}{},
	}

	vector := VectorArray{
		length:    0,
		array:     []interface{}{},
		capResize: 100,
	}

	factor := FactorArray{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}

	matrix := MatrixArray{
		length:       0,
		array:        make([][]interface{}, 1),
		lengthPart:   100,
		countOfPart:  1,
		factorResize: 2,
	}
	matrix.array[0] = make([]interface{}, matrix.lengthPart)

	inMiddleAdd(&simple, 100000, time.Now())
	inMiddleAdd(&vector, 100000, time.Now())
	inMiddleAdd(&factor, 100000, time.Now())
	inMiddleAdd(&matrix, 100000, time.Now())
}

func inMiddleAdd(arr Array, count int, t time.Time) {
	for i := 0; i < 10; i++ {
		arr.Add(i, i)
	}
	for i := 10; i < count; i++ {
		arr.Add(i, i/2)
	}
	outputString := ""
	switch arr.(type) {
	case *FactorArrayOld:
		outputString = "oldFactor"
	case *FactorArray:
		outputString = "newFactor"
	case *SimpleArrayOld:
		outputString = "oldSimple"
	case *SimpleArray:
		outputString = "newSimple"
	case *MatrixArray:
		outputString = "matrix"
	case *VectorArray:
		outputString = "vector"
	}
	outputString = fmt.Sprintf("%s finished, count of operations: %d, time in milliseconds: %d", outputString, count, time.Now().Sub(t).Milliseconds())
	fmt.Println(outputString)
}

func testAllArraysParallels() {
	fmt.Println("Test all arrays")
	simple := SimpleArray{
		length: 0,
		array:  []interface{}{},
	}

	vector := VectorArray{
		length:    0,
		array:     []interface{}{},
		capResize: 1000,
	}

	factor := FactorArray{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}

	matrix := MatrixArray{
		length:       0,
		array:        make([][]interface{}, 1),
		lengthPart:   1000,
		countOfPart:  1,
		factorResize: 2,
	}
	matrix.array[0] = make([]interface{}, matrix.lengthPart)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go AddManyElements(&simple, 100000, wg, time.Now())

	wg.Add(1)
	go AddManyElements(&vector, 100000, wg, time.Now())

	wg.Add(1)
	go AddManyElements(&factor, 100000, wg, time.Now())

	wg.Add(1)
	go AddManyElements(&matrix, 100000, wg, time.Now())

	wg.Add(1)
	go makeSlice(100000, wg, time.Now())

	wg.Add(1)
	go app(100000, wg, time.Now())

	wg.Wait()
}

func AddManyElements(arr Array, count int, wg *sync.WaitGroup, t time.Time) {
	for i := 0; i < count; i++ {
		arr.Add(i, i)
	}
	outputString := ""
	switch arr.(type) {
	case *SimpleArray:
		outputString = "simple"
	case *VectorArray:
		outputString = "vector"
	case *FactorArray:
		outputString = "factor"
	case *MatrixArray:
		outputString = "matrix"
	}
	outputString = fmt.Sprintf("%s finished, count of operations: %d, time in milliseconds: %d", outputString, count, time.Now().Sub(t).Milliseconds())
	fmt.Println(outputString)
	wg.Done()
}

// создание слайса сразу того размера который нужен, избигаем копирования
func makeSlice(count int, wg *sync.WaitGroup, t time.Time) {
	arr := make([]interface{}, count)
	for i := 0; i < count; i++ {
		arr[i] = i
	}
	fmt.Printf("make finished, count of operations: %d, time in milliseconds: %d\n", count, time.Now().Sub(t).Milliseconds())
	wg.Done()
}

// добавление элементов в стандартный динамический массив
func app(count int, wg *sync.WaitGroup, t time.Time) {
	arr := []interface{}{}
	for i := 0; i < count; i++ {
		arr = append(arr, i)
	}
	fmt.Printf("app finished, count of operations: %d, time in milliseconds: %d\n", count, time.Now().Sub(t).Milliseconds())
	wg.Done()
}

func testNewAndOldArrays() {
	fmt.Println("Test new and old arrays")
	newFactor := FactorArray{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	oldFactor := FactorArrayOld{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	inMiddleAdd(&newFactor, 100000, time.Now())
	inMiddleAdd(&oldFactor, 100000, time.Now())

	newFactor = FactorArray{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	oldFactor = FactorArrayOld{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	inMiddleAdd(&oldFactor, 100000, time.Now())
	inMiddleAdd(&newFactor, 100000, time.Now())

	newFactor = FactorArray{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	oldFactor = FactorArrayOld{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	inMiddleAdd(&oldFactor, 100000, time.Now())
	inMiddleAdd(&newFactor, 100000, time.Now())

	newFactor = FactorArray{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	oldFactor = FactorArrayOld{
		length:       0,
		array:        []interface{}{},
		factorResize: 2,
	}
	inMiddleAdd(&newFactor, 100000, time.Now())
	inMiddleAdd(&oldFactor, 100000, time.Now())

	oldSimple := SimpleArrayOld{
		length: 0,
		array:  []interface{}{},
	}
	newSimple := SimpleArray{
		length: 0,
		array:  []interface{}{},
	}
	inMiddleAdd(&newSimple, 50000, time.Now())
	inMiddleAdd(&oldSimple, 50000, time.Now())

	oldSimple = SimpleArrayOld{
		length: 0,
		array:  []interface{}{},
	}
	newSimple = SimpleArray{
		length: 0,
		array:  []interface{}{},
	}
	inMiddleAdd(&oldSimple, 50000, time.Now())
	inMiddleAdd(&newSimple, 50000, time.Now())

	oldSimple = SimpleArrayOld{
		length: 0,
		array:  []interface{}{},
	}
	newSimple = SimpleArray{
		length: 0,
		array:  []interface{}{},
	}
	inMiddleAdd(&oldSimple, 50000, time.Now())
	inMiddleAdd(&newSimple, 50000, time.Now())

	oldSimple = SimpleArrayOld{
		length: 0,
		array:  []interface{}{},
	}
	newSimple = SimpleArray{
		length: 0,
		array:  []interface{}{},
	}
	inMiddleAdd(&newSimple, 50000, time.Now())
	inMiddleAdd(&oldSimple, 50000, time.Now())
}
