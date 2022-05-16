package main

import (
	"fmt"
	"math"
)

func main() {

	change := 10
	M := []int{1, 2, 5, 10}
	r := dpCoinChange(change, M)
	fmt.Printf("%v coin(s) is/are needed to give change of %v with these denominations %+v\n", r, change, M)
}

func dpCoinChange(change int, M []int) int {

	rows := len(M)
	columns := change + 1
	H := generateMatrix(rows, columns)

	for r := 1; r < rows; r++ {
		for c := 1; c < columns; c++ {

			if c < M[r] {
				H[r][c] = H[r-1][c]
			} else if r == c {
				H[r][c] = 1
			} else {
				if r == 1 {
					H[r][c] = H[r][c-M[r]] + 1
				} else {
					H[r][c] = min(H[r][c-M[r]]+1, H[r-1][c])
				}

			}
		}
	}

	return H[rows-1][columns-1]

}

func min(i1, i2 int) int {
	return int(math.Min(float64(i1), float64(i2)))
}

func generateMatrix(rows, columns int) [][]int {
	H := make([][]int, rows)
	for i := 0; i < len(H); i++ {
		H[i] = make([]int, columns)
	}

	return H
}
