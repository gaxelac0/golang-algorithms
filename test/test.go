package main

import (
	"fmt"
)

func main() {

	r := pow(2, 4)
	fmt.Printf("%v", r)
}

func printMatrix() {

	var m [3][4]int

	fmt.Printf("%v\n", len(m))
	fmt.Printf("%v\n", len(m[0]))

}

func pow(a, n int) int {
	if n == 2 {
		return a * a
	} else {
		quad := pow(a, n/2)
		return quad * quad
	}
}
