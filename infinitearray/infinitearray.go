package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64
	x = 11
	fmt.Printf("Searching for: %v\n", x)

	arr := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	for i := 0; i < 10; i++ {
		arr = append(arr, math.Inf(1))
	}
	fmt.Printf("arr: = %+v\n", arr[0:10])

	e := findExp(arr, x)

	start := int(math.Pow(2, float64(e-1)))
	end := int(math.Pow(2, float64(e)))

	fmt.Printf("Sarching between %v and %v\n", start, end)

	pos := findPos(arr, start, end, x)
	fmt.Printf("Position: %v", pos)

}

/**
Binary search implementation to look for a position between a range of positions of the array.
In particular:
 - The array is ordered.
 - It n ordered elements but n is unknown.
 - After the n elements, +Inf are placed until infinitum.

Temporal cost: (log n)
**/
func findPos(arr []float64, start int, end int, x float64) int {

	half := (start + end) / 2
	isHalfInf := math.IsInf(arr[half], 1)
	isHalfplus1Inf := math.IsInf(arr[half+1], 1)

	if start > end {
		return -1
	} else if arr[half] == x {
		return half
	} else if arr[half] > x || isHalfInf {
		return findPos(arr, start, half, x)
	} else if arr[half] < x && !isHalfplus1Inf {
		return findPos(arr, half+1, end, x)
	} else {
		return -1
	}
}

/*
Returns a base two exponent.
That number (e) is used to calculate a range of numbers in which the position of the number x may be in the array.

x is contained in: {arr[e-1:e]}
x C {arr[e-1 : e]}

The temporal cost is logarithmic because this algorithm is not using a linear expresion to look for the index.
Instead is using s base two exponents to delimit the range.

Temporal cost: O(log n)
**/
func findExp(arr []float64, x float64) int {

	i := 0
	for idx := 0; arr[idx] < x && arr[idx] != math.Inf(1); i++ {
		idx = int(math.Pow(2, float64(i))) // 2^i
	}

	return i - 1
}
