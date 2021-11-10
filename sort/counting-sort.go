package sort

// CountingSort - sorting algorithm
// https://en.wikipedia.org/wiki/Counting_sort
func CountingSort(data []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	sortedIndex := 0
	length := len(data)

	for i := 0; i < length; i++ {
		bucket[data[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			data[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return data
}
