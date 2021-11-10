package sort

// BubbleSort - sorting algorithm
// https://en.wikipedia.org/wiki/Bubble_sort
func BubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}

	return data
}
