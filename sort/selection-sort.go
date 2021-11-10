package sort

// SelectionSort - sorting algorithm
// https://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(data []int) []int {
	length := len(data)

	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[length-i-1], data[maxIndex] = data[maxIndex], data[length-i-1]
	}

	return data
}
