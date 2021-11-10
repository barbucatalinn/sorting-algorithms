package sort

// QuickSort - sorting algorithm
// https://en.wikipedia.org/wiki/Quicksort
func QuickSort(data []int, low, high int) []int {
	if low < high {
		var p int
		data, p = partition(data, low, high)
		data = QuickSort(data, low, p-1)
		data = QuickSort(data, p+1, high)
	}

	return data
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
