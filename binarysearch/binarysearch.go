package main

import "fmt"

/**
Binary Search implementation without indexes.
The array passed by MUST be ordered.
Temporal Cost: O(log n)
**/
func Bs(arr []int, elem int) bool {
	half := (len(arr) - 1) / 2

	if len(arr) == 0 {
		return false
	}
	if arr[half] == elem {
		return true
	} else if arr[half] > elem {
		return Bs(arr[0:half], elem)
	} else {
		return Bs(arr[half+1:len(arr)-1], elem)
	}
}

/**
Binay Search implementation with indexes.
The array passed by args MUST be ordered.
Temporal Cost: O(log n)
**/
func BsIdx(arr []int, start int, end int, elem int) bool {

	if start == end {
		return (arr[start] == elem)
	} else if start > end {
		return false
	}

	mitad := (end + start - 1) / 2
	if arr[mitad] == elem {
		return true
	} else if arr[mitad] > elem {
		return BsIdx(arr, start, mitad, elem)
	} else {
		return BsIdx(arr, mitad+1, end, elem)
	}
}

func main() {

	elem := 5
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	bPresent := Bs(arr, elem)
	bPresentIdx := BsIdx(arr, 0, len(arr)-1, elem)

	fmt.Printf("[without index] %d is present in arr? %v\n", elem, bPresent)
	fmt.Printf("[with index] %d is present in arr? %v\n", elem, bPresentIdx)
	fmt.Printf("Equal results? %v", (bPresent == bPresentIdx))
}
