package main

import (
	"fmt"
	"math/rand"
	"time"
	sortfunctions "sortAlgoritms/sortFunctions"
)

func main() {
	arr := createArr()
	fmt.Println(arr[:10])
	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	// Пузырьком
	start := time.Now()	
	sortfunctions.BubleSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Пузырьковая сортировака O(n^2) занимает: ", duration)
	// Выбором
	start = time.Now()	
	sortfunctions.SelectionSort(emptyArr)
	duration = time.Since(start)
	fmt.Println("Сортировка выбором / двоичный поиск O(n^2) занимает: ", duration)
	// Вставкой
	start = time.Now()	
	sortfunctions.InsertionSort(emptyArr)
	duration = time.Since(start)
	fmt.Println("Сортировка вставкой O(n^2) занимает: ", duration)
	// Быстрая сортировка
	start = time.Now()	
	sortfunctions.QuickSort(emptyArr)
	duration = time.Since(start)
	fmt.Println("Быстрая сортировка O(n x log n) занимает: ", duration)
	// Сортировка слиянием
	start = time.Now()	
	sortfunctions.MergeSort(emptyArr)
	duration = time.Since(start)
	fmt.Println("Сортировка слиянием O(n x log n) занимает: ", duration)
}

func createArr() (arr []int) {
	sizeArr := 50000
	arr = make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
	}
	return
}
