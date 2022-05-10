package main

import "fmt"

func main() {

	V := []int{0, 1, 3, 2, 5, 1, 0}
	peaks := bbPeaks(V, 0, len(V)-1)

	fmt.Printf("Peaks: %v\n", peaks)

}

func bbPeaks(V []int, start, end int) int {

	if V[start] >= V[start+1] {
		return V[start]
	}

	if V[end] >= V[end-1] {
		return V[end]
	}

	half := (start + end) / 2
	if V[half] >= V[half-1] && V[half] >= V[half+1] {
		return V[half]
	} else if V[half] < V[half+1] {
		return bbPeaks(V, half, end)
	} else {
		return bbPeaks(V, start, half)
	}

}
