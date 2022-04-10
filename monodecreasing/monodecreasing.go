package main

import (
	"fmt"
	"math"
)

// Random function
type a_function func(int) int

/**
findRoot is a binary search algorithm implementation to find the
position in which an arbitrary monotonous-decreasing function becomes zero.
Returns -1 if theres no null value for the function
**/
func findRoot(f a_function, start int, end int) int {
	half := (end + start) / 2
	if start > end {
		return -1
	}
	if f(half) == 0 {
		return half
	} else if f(half) < 0 {
		return findRoot(f, start, half)
	} else {
		return findRoot(f, half, end)
	}
}

/**
Funcion that finds the minor element in the x axis, using exponents of two.
In which value, the function already became zero in the y axis.
Temporal Cost: O(1)
**/
func findExp(f a_function) int {

	i := 0
	for idx := 0; f(idx) >= 0; i++ {
		idx = int(math.Pow(2, float64(i))) // 2^i
	}

	return i
}

func main() {

	f := mock3
	e := findExp(f)

	start := int(math.Pow(2, float64(e-2)))
	end := int(math.Pow(2, float64(e-1)))

	fmt.Printf("Sarching between %v and %v\n", start, end)

	// Apply binary search between 2^(e-1) 2^(e) to find x:{xeR/f(x)=0)
	zeroval := findRoot(f, start, end)
	fmt.Printf("f(%v) = 0", zeroval)
}

/**
Random mocked arbitrary function that is monotonous-decreasing.
Can be replaced for any other function that meet the conditions.
**/
func mock(x int) int {
	if x == 1637 {
		return 0
	} else if x < 1637 {
		return 1
	} else {
		return -1
	}
}

/**
Random mocked arbitrary function that is monotonous-decreasing.
Can be replaced for any other function that meet the conditions.
**/
func mock2(x int) int {
	if x == 115 {
		return 0
	} else if x < 115 {
		return 1
	} else {
		return -1
	}
}

/**
Random mocked arbitrary function that is monotonous-decreasing.
Can be replaced for any other function that meet the conditions.
**/
func mock3(x int) int {
	if x == 2048 {
		return 0
	} else if x < 2048 {
		return 1
	} else {
		return -1
	}
}

/**
Random mocked arbitrary function that is monotonous-decreasing.
Can be replaced for any other function that meet the conditions.
**/
func mock4(x int) int {
	if x == 0 {
		return 0
	} else if x < 0 {
		return 1
	} else {
		return -1
	}
}

/**
Random mocked arbitrary function that is monotonous-decreasing.
Can be replaced for any other function that meet the conditions.
**/
func mock5(x int) int {
	if x == -1 {
		return 0
	} else if x < -1 {
		return 1
	} else {
		return -1
	}
}
