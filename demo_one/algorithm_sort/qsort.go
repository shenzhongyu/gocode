package algorithm_sort
// 快速排序
func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort(values, 0, len(values) - 1)
}
// [11, 43, 2554, 87, 8]
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {

		for j >= p && values[j] >= temp {
			j --
		}

		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i ++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}

		values[p] = temp
		if p - left > 1 {
			quickSort(values, left, p - 1)
		}

		if right - p > 1 {
			quickSort(values, p + 1, right)
		}
	}
}