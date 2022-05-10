package main

import "fmt"

func main() {

	n := 7
	r := recursiveFibonacci(n)
	fmt.Printf("Fibonacci de %v es %v", n, r)
}

func recursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
	}
}

func dpFibonacci() {

}
