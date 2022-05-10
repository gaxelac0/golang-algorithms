package main

import (
	"fmt"
	"math"
)

/**
Implementation of calculation Entire Square Root of a number.
It uses binary search algorithm in the background for a O(log n) temporal cost.
Filing = Radicando
**/
func integerSquareRoot(start int, end int, filing int) int {
	half := (end + start - 1) / 2
	sqrHalf := int(math.Pow(float64(half), 2))
	sqrHalfPlusOne := int(math.Pow(float64(half+1), 2))

	fmt.Printf("sqrHalf: %v\nsqrHalfPlusOne: %v\n", sqrHalf, sqrHalfPlusOne)

	if start > end {
		return -1
	} else if filing == 0 || filing == 1 {
		return filing
	} else if sqrHalf <= filing && sqrHalfPlusOne > filing {
		return half
	} else if half > filing {
		return integerSquareRoot(half+1, end, filing)
	} else {
		return integerSquareRoot(start, half, filing)
	}
}

func sqrt(filing int) int {
	return integerSquareRoot(1, filing, filing)
}

func main() {
	radicando := 4
	raiz := sqrt(radicando)
	fmt.Printf("raiz cuadrada entera de %v: %v", radicando, raiz)
}
