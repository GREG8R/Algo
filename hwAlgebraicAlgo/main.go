//НОД, Степень, Простые числа, Числа Фибоначчи.
//Цель: Написать и сравнить разные алгоритмы нахождения наибольшего общего делителя,
// возведения числа в целую степень, поиска простых чисел и вычисления чисел фибоначчи.
//Написать следующие алгоритмы и сравнить их быстродействие и результаты.
//Приложить скриншоты/ссылки на сравнительную таблицу проведённых экспериментов.
//
//На выбор: 1 или 2 задание, а также 3 или 4.
//
//1. Алгоритм Евклида поиска НОД макс. 4 байта
//1a. Через вычитание
//+1 байт 1b. Через остаток
//+1 байт 1c. Через битовые операции
//+2 байт Составить сравнительную таблицу времени работы алгоритмов для разных начальных данных.
//Написать, какие пункты выполнены и сколько времени ушло на выполнение домашнего задания.
//
//2. Алгоритм возведения в степень макс. 4 байта
//2а. Итеративный (n умножений)
//+1 байт 2b. Через степень двойки с домножением
//+1 байт 2c. Через двоичное разложение показателя степени.
//+2 байт Составить сравнительную таблицу времени работы алгоритмов для разных начальных данных.
//Написать, какие пункты выполнены и сколько времени ушло на выполнение домашнего задания.
//
//3. Алгоритмы поиска кол-ва простых чисел до N макс. 6 байт
//3a. Через перебор делителей.
//+1 байт 3b. Несколько оптимизаций перебора делителей, с использованием массива.
//+1 байт 3c. Решето Эратосфена со сложностью O(n log log n).
//+1 байт 3d. Решето Эратосфена с оптимизацией памяти: битовая матрица, по 32 значения в одном int
//+1 байт 3e. Решето Эратосфена со сложностью O(n)
//+2 байт Составить сравнительную таблицу времени работы алгоритмов для разных начальных данных.
//Написать, какие пункты выполнены и сколько времени ушло на выполнение домашнего задания.
//
//4. Алгоритм поиска чисел Фибоначчи макс. 6 байта
//4a. Через рекурсию
//+1 байт 4b. Через итерацию
//+1 байт 4c. По формуле золотого сечения
//+2 байт 4d. Через умножение матриц
//+2 байт Составить сравнительную таблицу времени работы алгоритмов для разных начальных данных.
package main

import (
	"fmt"
	"time"
)

func main(){
	fibLaunch()
	powerLaunch()
}

func fibLaunch(){
	n := 45
	fmt.Printf("fibonacci test n = %d\n", n)

	t := time.Now()
	result := FibRecursive(n)
	fmt.Printf("fib recursive result: %d, time: %d\n", result, time.Now().Sub(t).Nanoseconds())

	t = time.Now()
	result = FibIteration(n)
	fmt.Printf("fib iteration result: %d, time: %d\n", result, time.Now().Sub(t).Nanoseconds())

	t = time.Now()
	result = FibFormula(n)
	fmt.Printf("fib formula result: %d, time: %d\n", result, time.Now().Sub(t).Nanoseconds())

	t = time.Now()
	result = FibMatrix(n)
	fmt.Printf("fib matrix result: %d, time: %d\n\n", result, time.Now().Sub(t).Nanoseconds())
}

func powerLaunch(){
	number := 1.0000000001
	pow := 10000000000

	fmt.Printf("power test number = %f, power = %d\n", number, pow)

	t := time.Now()
	result := PowOfPowTwo(number, pow)
	fmt.Printf("power of power two result: %f, time: %d\n", result, time.Now().Sub(t).Milliseconds())

	t = time.Now()
	result = PowOfComposition(number, pow)
	fmt.Printf("power of composition result: %f, time: %d\n", result, time.Now().Sub(t).Milliseconds())

	t = time.Now()
	result = FastPow(number, pow)
	fmt.Printf("fast power result: %f, time: %d\n\n", result, time.Now().Sub(t).Milliseconds())

	number = 2
	pow = 20
	fmt.Printf("test number = %f, power = %d\n", number, pow)
	t = time.Now()
	result = PowOfPowTwo(number, pow)
	fmt.Printf("power of power two result: %f, time: %d\n", result, time.Now().Sub(t).Milliseconds())

	t = time.Now()
	result = PowOfComposition(number, pow)
	fmt.Printf("power of composition result: %f, time: %d\n", result, time.Now().Sub(t).Milliseconds())

	t = time.Now()
	result = FastPow(number, pow)
	fmt.Printf("fast power result: %f, time: %d\n\n", result, time.Now().Sub(t).Milliseconds())
}