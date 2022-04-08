package main

import (
	"fmt"
	"math"
)

/**
findNullValueForF is a binary search algorithm implementation to find the
position in which an arbitrary monotonous-decreasing function becomes zero.
**/
func findNullValueForF(start int, end int) int {
	half := (end + start - 1) / 2
	if start > end {
		return -1
	}
	if f(half) == 0 {
		return half
	} else if f(half) < 0 {
		return findNullValueForF(start, half)
	} else {
		return findNullValueForF(half, end)
	}
}

/**
Funcion that finds the minor element in the x axis, using exponents of two.
In which value, the function already became zero in the y axis.
Temporal Cost: O(1)
**/
func findExp() int {

	i := 0
	for idx := 0; f(idx) >= 0; i++ {
		idx = int(math.Pow(2, float64(i)))
	}

	return i
}

func main() {
	e := findExp()

	start := int(math.Pow(2, float64(e-2)))
	end := int(math.Pow(2, float64(e-1)))
	// Apply binary search between 2^(e-1) 2^(e) to find x:{xeR/f(x)=0)
	zeroval := findNullValueForF(start, end)
	fmt.Printf("f(%v) = 0", zeroval)
}

/**
Random mocked arbitrary function that is monotonous-decreasing.
Can be replaced for any other function that meet the conditions.
**/
func f(x int) int {
	if x == 1637 {
		return 0
	} else if x < 1637 {
		return 1
	} else {
		return -1
	}
}
