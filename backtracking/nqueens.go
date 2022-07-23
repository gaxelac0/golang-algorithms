package main

type Matrix struct {
	m [][]int
}

func main() {

	n := 8
	mtx := getMatrix(n)

	queensProblem(mtx, 0)
	// mtx.displayMatrix()

}

func queensProblem(mtx Matrix, stage int) {

	for c := 0; c < len(mtx.m[0]); c++ {

		if stage == len(mtx.m) {
			mtx.displayMatrix()
			return
		}

		mtx.m[stage][c] = 1
		if mtx.isValid(stage, c) {
			queensProblem(mtx, stage+1)
		}
		mtx.m[stage][c] = 0
	}
}

func getMatrix(n int) Matrix {

	var mtx Matrix

	m := make([][]int, n)
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, n)
	}

	mtx.m = m
	return mtx
}

func (mtx Matrix) isValid(r, c int) bool {
	// check in rows
	for _r := 0; _r < len(mtx.m); _r++ {
		if _r != r && mtx.m[_r][c] == 1 {
			return false
		}
	}

	// check in columns
	for _c := 0; _c < len(mtx.m[0]); _c++ {
		if _c != c && mtx.m[r][_c] == 1 {
			return false
		}
	}

	// check in diagonals
	for _r := 0; _r < len(mtx.m); _r++ {
		for _c := 0; _c < len(mtx.m[0]); _c++ {
			if _r == r && _c == c {
				continue
			}
			if mtx.m[_r][_c] == 1 && (_r+_c == r+c || _r-_c == r-c) {
				return false
			}
		}
	}

	return true
}

func (mtx Matrix) displayMatrix() {
	for r := 0; r < len(mtx.m); r++ {
		for c := 0; c < len(mtx.m[r]); c++ {

			print(mtx.m[r][c])
			if c == len(mtx.m)-1 {
				println()
			}
		}
	}
}
