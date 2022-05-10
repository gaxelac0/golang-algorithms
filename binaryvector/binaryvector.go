package main

import "fmt"

func main() {
	V := []int{0, 0, 0, 0, 0, 0, 1, 1}
	c := countOnes(V, 0, len(V)-1)
	fmt.Printf("Ones quantity: %v\n", c)
}

func countOnes(V []int, start, end int) int {

	half := (start + end) / 2
	if V[half] == 1 {

		if V[half-1] == 0 {
			return end + 1 - half
		} else {
			return (end + 1 - half) + countOnes(V, start, half-1)
		}
	} else {
		if V[half+1] == 1 {
			return end - half
		} else {
			return 0 + countOnes(V, half+1, end)
		}

	}

}
