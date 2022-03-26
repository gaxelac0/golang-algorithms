package main

import (
	"fmt"
)

/*
Iterative implementation of the bubblesort algorithm.
The temporal complexity of the method is O(n²)
*/
func bubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	wanted := []int{3, 7, 12, 111}

	list := []int{7, 12, 111, 3}
	sorted := bubbleSort(list)

	out := fmt.Sprintf("sorted %v, want %v", sorted, wanted)
	fmt.Println(out)
}
