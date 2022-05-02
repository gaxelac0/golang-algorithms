package main

import (
	"fmt"
	"math"
)

func main() {

	x := []string{"", "A", "B", "A", "C", "A", "D", "A", "C"}
	y := []string{"", "A", "B", "C", "B", "D", "A"}
	lcs := longestCommonSubseq(x, y) // A, B, C, D, A

	fmt.Printf("%v", lcs)
}

func longestCommonSubseq(x []string, y []string) []string {

	m := generationDP(x, y)
	lcs := reconstructionDP(x, y, m)

	return lcs

}

// TODO need fixing
func reconstructionDP(x []string, y []string, m [][]int) []string {

	//i := 0
	result := make([]string, 0) // A, B, C, D, A
	for r := len(m) - 1; r >= 1; {
		for c := len(m[0]) - 1; c >= 1; {

			if x[r] == y[c] {
				result = append(result, x[r])
				r--
				c--
			} else {

				if m[r-1][c] == m[r][c-1] {
					c--
				} else {
					r--
				}
			}
		}
	}
	reverse(result)
	return result
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func generationDP(x []string, y []string) [][]int {

	m := createMatrix(len(x), len(y))

	for r := 0; r < len(x); r++ {
		for c := 0; c < len(y); c++ {

			if r == 0 || c == 0 {
				m[r][c] = 0
			} else {
				if x[r] == y[c] {
					m[r][c] = m[r-1][c-1] + 1
				} else {
					m[r][c] = int(math.Max(float64(m[r-1][c]), float64(m[r][c-1])))
				}
			}
		}
	}

	return m
}

func createMatrix(x int, y int) [][]int {

	m := make([][]int, x)

	for i := 0; i < x; i++ {
		m[i] = make([]int, y)
	}
	return m
}
