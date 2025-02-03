package binarysearch

import (
	"fmt"
	"math/rand/v2"
)

func binarySearch(needle int, arr []int) (bool, int, int) {
	low := 0
	high := len(arr)

	for low < high {
		mid := low + (high-low)/2
		value := arr[mid]

		if value == needle {
			return true, mid, value
		}

		if needle > value {
			low = mid + 1
		} else {
			high = mid
		}

	}

	return false, -1, -1
}

func bubbleSort(arr []int) []int {
	out := make([]int, len(arr))

	l := len(arr)
	for l != 0 {

		for i, v := range arr {
			if i+1 >= l {
				continue
			}

			if v > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], v
			}
		}

		out[l-1] = arr[l-1]
		arr = arr[:l-1]
		l = len(arr)
	}

	return out
}

func qsPartition(arr []int, low int, high int) ([]int, int) {
	pivot := arr[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	idx++
	arr[idx], arr[high] = arr[high], arr[idx]
	return arr, idx
}

func qs(arr []int, low int, high int) []int {
	if low >= high {
		return arr
	}

	var pivotIdx int
	arr, pivotIdx = qsPartition(arr, low, high)
	arr = qs(arr, low, pivotIdx-1)
	arr = qs(arr, pivotIdx+1, high)
	return arr
}

func quicksort(arr []int) []int {
	return qs(arr, 0, len(arr)-1)
}

func TestStructBinarySearch() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.IntN(10)
	}

	sort := quicksort(arr)
	fmt.Printf("size: %d \n", len(sort))
	fmt.Println(sort)
}
