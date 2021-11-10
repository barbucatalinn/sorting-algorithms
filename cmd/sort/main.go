package main

import (
	"fmt"

	"github.com/barbucatalinn/sorting-algorithms/sort"
)

func main() {
	unsorted := []int{9, 4, 7, 3, 1, 8, 10, 2, 5, 6, 0}

	bubbleSortSorted := sort.BubbleSort(unsorted)
	bucketSortSorted := sort.BucketSort(unsorted, len(unsorted))
	countingSortSorted := sort.CountingSort(unsorted, len(unsorted))
	heapSortSorted := sort.HeapSort(unsorted)
	insertionSortSorted := sort.InsertionSort(unsorted)
	mergeSortSorted := sort.MergeSort(unsorted)
	quickSortSorted := sort.QuickSort(unsorted, 0, len(unsorted)-1)
	radixSortSorted := sort.RadixSort(unsorted)
	randomQuickSortSorted := sort.RandomQuickSort(unsorted)
	selectionSortSorted := sort.SelectionSort(unsorted)
	shellSortSorted := sort.ShellSort(unsorted)

	fmt.Println("Unsorted:", unsorted)

	fmt.Println("Bubble Sort:", bubbleSortSorted)
	fmt.Println("Bucket Sort:", bucketSortSorted)
	fmt.Println("Counting Sort:", countingSortSorted)
	fmt.Println("Heap Sort:", heapSortSorted)
	fmt.Println("Insertion Sort:", insertionSortSorted)
	fmt.Println("Merge Sort:", mergeSortSorted)
	fmt.Println("Quick Sort:", quickSortSorted)
	fmt.Println("Radix Sort:", radixSortSorted)
	fmt.Println("Random Quick Sort:", randomQuickSortSorted)
	fmt.Println("Selection Sort:", selectionSortSorted)
	fmt.Println("Shell Sort:", shellSortSorted)
}
