package sort

// InsertionSort - sorting algorithm
// https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			j := i - 1
			temp := data[i]

			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}

			data[j+1] = temp
		}
	}

	return data
}
