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
"Exponential Search" / "Busqueda Exponencial"
Funcion that finds the minor element in the x axis, using exponents of two.
In which value, the function already became zero in the y axis.

Returns a base two exponent.
That number (e) is used to calculate a range of numbers in which the position of the number x may be in the array.

x is contained in: {arr[e-1:e]}
x C {arr[e-1 : e]}

The temporal cost is logarithmic because this algorithm is not using a linear expresion to look for the index.
Instead is using s base two exponents to delimit the range.
Temporal Cost: O(log n)
**/
func exponentialSearch(f a_function) int {

	i := 0
	for idx := 0; f(idx) >= 0; i++ {
		idx = int(math.Pow(2, float64(i))) // 2^i
	}

	return i - 1
}

func main() {

	f := mock
	e := exponentialSearch(f)

	start := int(math.Pow(2, float64(e-1)))
	end := int(math.Pow(2, float64(e)))

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
