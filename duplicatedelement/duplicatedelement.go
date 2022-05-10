package main

import "fmt"

func main2() {

	P := []int{0, 1, 2, 3, 4, 5, 4, 6}
	r := Bs_DuplicatedElement(P, 0, len(P)-1)
	fmt.Printf("Bad position number idx := %v\n", r)
}

/**
Binary Search impleementation with indexes.
**/
func Bs_DuplicatedElement(P []int, ini, end int) int {
	mid := (ini + end) / 2

	if P[mid] != mid {
		return mid
	}

	if P[mid] == mid {
		return Bs_DuplicatedElement(P, mid, end)
	} else {
		return Bs_DuplicatedElement(P, 0, mid)
	}
}
