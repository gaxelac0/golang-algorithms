package main

/*
Recursive implementation of MergeSort algorithm.
MergeSort splits the array in half recursively until the atomic case (one element), then merge in the correct order
the last two atomic elements into an auxiliar array called W. W is used to accomodate the original array later.
This last step is repeated until the original array is ordered.
Temporal complexity: O(n^1 log n) = O(n log n)
a=2
b=2
k=1
O(n^k log n)

a = b^k
2 = 2
*/
func rMergeSort(arr []int, start int, end int) {

	if start < end {
		mid := (start + end) / 2 // constant
		rMergeSort(arr, start, mid)
		rMergeSort(arr, mid+1, end)
		Merge(arr, start, end) // O(n)
	}
}

func MergeSort(arr []int, start int, end int) {
	rMergeSort(arr, start, end)
}

// Temporal Complexity: O(n)
func Merge(arr []int, start, end int) {

	w := make([]int, end-start+1) // or.. var w [end-start]int

	mid := (start + end) / 2
	i := start
	j := mid + 1

	for k := 0; k <= end-start; k++ {
		if j > end || arr[i] <= arr[j] && i < mid+1 {
			w[k] = arr[i]
			i++
		} else {
			w[k] = arr[j]
			j++
		}
	}

	for k := 0; k <= end-start; k++ {
		arr[start+k] = w[k]
	}
}
