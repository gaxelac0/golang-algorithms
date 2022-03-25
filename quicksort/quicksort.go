package main

import "fmt"

func main() {
	arr := []int{32, 64, 22, 11, 2, 36}
	wanted := []int{2, 11, 22, 32, 36, 64}

	QuickSort(arr, 0, len(arr)-1)

	fmt.Printf("sorted %v, wanted %v", arr, wanted)
}

func QuickSort(arr []int, start int, end int) {

	if start < end {
		pivot := pivot(arr, start, end)
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
}

func pivot(arr []int, start int, end int) int {

	p := arr[start]
	leftIdx := start + 1
	rightIdx := end

	for arr[leftIdx] <= p && leftIdx < end {
		leftIdx++
	}

	for arr[rightIdx] > p {
		rightIdx--
	}

	for leftIdx < rightIdx {
		aux := arr[rightIdx]
		arr[rightIdx] = arr[leftIdx]
		arr[leftIdx] = aux

		for arr[leftIdx] <= p && leftIdx < end {
			leftIdx++
		}

		for arr[rightIdx] > p {
			rightIdx--
		}
	}
	arr[start] = arr[rightIdx]
	arr[rightIdx] = p

	return rightIdx
}
