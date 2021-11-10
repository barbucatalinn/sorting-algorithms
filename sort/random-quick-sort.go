package sort

import "math/rand"

// RandomQuickSort - sorting algorithm
// It implements QuickSort using random pivoting
// https://en.wikipedia.org/wiki/Quicksort
func RandomQuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1
	pivot := rand.Int() % len(data)

	data[pivot], data[right] = data[right], data[pivot]

	for i, _ := range data {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}

	data[left], data[right] = data[right], data[left]

	RandomQuickSort(data[:left])
	RandomQuickSort(data[left+1:])

	return data
}
