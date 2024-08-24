package sortfunctions

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] { // сравниваем элемент из левой и из правой половины
			merged = append(merged, left[0]) // если из left меньше добавляем его
			left = left[1:]
		} else {
			merged = append(merged, right[0]) // иначе добавляем из right
			right = right[1:]
		}
	}
	merged = append(merged, left...) // склеиваем результаты, части left и right
	merged = append(merged, right...)
	return merged
}