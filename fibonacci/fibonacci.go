package main

import "fmt"

func main() {

	n := 7
	r := recursiveFibonacci(n)
	fmt.Printf("Fibonacci(R) de %v es %v\n", n, r)

	dp := dpFibonacci(n)
	fmt.Printf("Fibonacci(DP) de %v es %v\n", n, dp)
}

/**
Recursive implementation of Fibonacci sequence.
It is not performance because its calculations are not independant and are recalculated exponentially.
O(2^n)

a: 2
b: 1
k = 0

O(2^n)
**/
func recursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
	}
}

/*
Implementation using dynamic programming, it's performant because it doesn't recalculate the terms, instead
uses an array to refer to previously calculated values.
*/
func dpFibonacci(n int) int {

	T := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i < 2 {
			T[i] = i
		} else {
			T[i] = T[i-1] + T[i-2]
		}
	}

	return T[n]
}
