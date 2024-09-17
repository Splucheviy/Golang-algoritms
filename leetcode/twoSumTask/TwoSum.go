package twosumtask

import (
	"fmt"
)

var nums = []int{2, 7, 11, 15}
var target = 9

// пробедались по всему слайсу два раза, перебрали все значения, список, т.е. O(n)
func TwoSumFirstCase(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target-nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// идея в хэш-таблице, мапе, работаем в две итерации, первой составляем мапу, второй ищем targets-значение в мапе,
// если есть совпадение - запоминаем и выводим индексы, так как мапа - то O(1)
func TwoSumSecondCase(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		delta := target - num
		if j, ok := m[delta]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}

// один проход
func TwoSumThirdCase(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		delta := target - num
		if j, ok := m[delta]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return []int{}
}

func TwoSumFirstCaseRun() {
	fmt.Println(TwoSumFirstCase(nums, target))
}

func TwoSumSecondCaseRun() {
	fmt.Println(TwoSumSecondCase(nums, target))
}

func TwoSumThirdCaseRun() {
	fmt.Println(TwoSumThirdCase(nums, target))
}
