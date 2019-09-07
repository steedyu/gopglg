package qsort

func quickSort(values []int, left int, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i < j {
		for j >= p && temp <= values[j] {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && temp >= values[i] {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
