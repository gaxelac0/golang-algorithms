package main

import (
	"fmt"
	"math"
)

func main() {
	SnarkHunting()
}

/**
Based on Lewis Carroll surrealist poems.
An Explorer is looking for Snarks but Boojums make people dissapear from the earth.

Exercise to resolve by using Dynamic Programming.

Having a nxn 2-d slice that shows 1 when the position has an Snark,
 -2 when it has a Boojum and 0 when it doesn't have neither and Snark nor an Boojum.

 What is the maximum quantity of Snark we can found without meeting a Boojum having in count:
  - Explorer can only move to the right or down.
**/
func SnarkHunting() {

	// base static 2d slice
	M := obtainOriginalMatrix()

	// generate dp matrix to consume
	// the result of maximum quantity of Snark found is in the last position of the 2D slice
	H := generateDP(M)
	fmt.Printf("Snark hunted: %v\n", H[len(H)-1][len(H[0])-1].sum)

	// for the trayectory we use the route from the last position saving the coords.
	trayectory := reconstructionDP(H)
	fmt.Printf("Trayectory coordenates: %+v\n", trayectory)
}

func reconstructionDP(H [][]obj) []coord {
	var coords []coord = make([]coord, 0)
	for r := len(H) - 1; r > 1; {
		for c := len(H[0]) - 1; c > 1; {

			if H[r][c].v == 1 {
				coords = append(coords, *NewCoord(r, c))
			}

			// Los valores de los dos precedentes son 0 o -1
			if (H[r-1][c].v == -1 || H[r-1][c].v == 0) && (H[r][c-1].v == -1 || H[r][c-1].v == 0) {

				// si las sumas son iguales, voy a la iqzuierda c--
				// si el valor de rriba es mas grande, también c--
				// si el valor es menor, voy para arriba
				if H[r-1][c].sum == H[r][c-1].sum { // can unify this ***
					c--
				} else {

					if H[r-1][c].sum > H[r][c-1].sum {
						r--
					} else { // can unify this *** <=
						c--
					}
				}
			} else if H[r-1][c].v == -1 || H[r-1][c].v == 0 {
				c--
			} else {
				r--
			}

		}
	}

	return reverse(coords)
}

func generateDP(M [][]int) [][]obj {

	H := make([][]obj, 9)
	for i := range H {
		H[i] = make([]obj, 9)
	}

	for r := 0; r < 9; r++ {

		for c := 0; c < 9; c++ {
			H[r][c] = *NewObj(0, 0)
		}
	}

	for r := 1; r < 9; r++ {

		for c := 1; c < 9; c++ {

			// The two precedents are -1
			if H[r-1][c].v == -1 && H[r][c-1].v == -1 {
				H[r][c].v = -1
				H[r][c].sum = max(H[r-1][c].sum, H[r][c-1].sum)
			} else if M[r-1][c-1] == 0 { // neither boojum nor snark
				H[r][c].v = 0
				H[r][c].sum = max(H[r-1][c].sum, H[r][c-1].sum)
			} else if M[r-1][c-1] == 1 { // snark
				H[r][c].v = 1
				H[r][c].sum = max(H[r-1][c].sum, H[r][c-1].sum) + 1
			} else { // is a boojum
				H[r][c].v = -1
				H[r][c].sum = max(H[r-1][c].sum, H[r][c-1].sum)
			}

		}
	}

	return H
}

func max(i1, i2 int) int {
	return int(math.Max(float64(i1), float64(i2)))
}

func obtainOriginalMatrix() [][]int {

	M := make([][]int, 8)
	for i := range M {
		M[i] = make([]int, 8)
	}

	for r := 0; r < 8; r++ {

		for c := 0; c < 8; c++ {
			M[r][c] = 0
		}
	}

	M[0][1] = 1
	M[0][3] = 1
	M[0][4] = 1
	M[2][1] = 1
	M[2][4] = 1
	M[3][2] = 1
	M[3][7] = 1
	M[4][4] = 1
	M[6][0] = 1
	M[6][4] = 1
	M[6][6] = 1
	M[7][2] = 1

	M[0][5] = 2
	M[0][7] = 2
	M[1][3] = 2
	M[2][7] = 2
	M[3][0] = 2
	M[3][3] = 2
	M[3][6] = 2
	M[5][1] = 2
	M[5][3] = 2
	M[5][5] = 2
	M[6][7] = 2

	return M
}

type obj struct {
	v, sum int
}

type coord struct {
	r, c int
}

func NewObj(v, sum int) *obj {
	return &obj{v, sum}
}

func NewCoord(r, c int) *coord {
	return &coord{r: r, c: c}
}
func reverse(ss []coord) []coord {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}

	return ss
}
