package sortfunctions

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0] // выбираем опорный элемент
	var less, greater []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(QuickSort(less), pivot) // элементы меньше опорного помещаем до
	result = append(result, QuickSort(greater)...) // элементы больше опорного помещаем после
	return result
}