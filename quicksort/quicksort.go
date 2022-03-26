package main

import "fmt"

func main() {
	arr := []int{32, 64, 22, 11, 2, 36}
	wanted := []int{2, 11, 22, 32, 36, 64}

	QuickSort(arr, 0, len(arr)-1)

	fmt.Printf("sorted %v, wanted %v", arr, wanted)
}

/*
Recursive implementation of QuickSort algorithm.
QuickSort makes use of a Pivot element, which is the first element of the array sent by parameter in this implementation.
All the rest of elements are ordered at the sides of the Pivot, those smaller than the pivot go to the left, otherwise to the right of the pivot.
This process is done recursively until all the sub-arrays are ordered.
Temporal complexity: O(n^(log2 2)) = O(n)
a=2
b=2
k=0
O(n^(logb a))

a > b^k
2 > 1
*/
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
