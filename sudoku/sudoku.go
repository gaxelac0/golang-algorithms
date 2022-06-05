package main

import "fmt"

type Sudoku struct {
	matrix [][]int
}

func main() {
	sudoku := generateSudoku()
	bSolution := sudoku.solve()
	fmt.Printf("Tiene solución?: %v\n\n", bSolution)
	if bSolution {
		printMatrix(sudoku.matrix)
	}
}

func (s Sudoku) solve() bool {

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			if s.matrix[r][c] == 0 {
				for v := 1; v <= 9; v++ {

					if s.isValid(r, c, v) {
						s.setVal(r, c, v)
						if s.solve() {
							return true
						} else {
							// backtrack
							s.setVal(r, c, 0)
						}
					}
				}
				return false
			}
		}
	}

	return true
}

func (s Sudoku) isValid(r, c, v int) bool {

	// validate the number is not present in the column
	for act_row := 0; act_row < 9; act_row++ {
		if s.matrix[act_row][c] == v {
			return false
		}
	}

	// validate the number is not present in the row
	for act_col := 0; act_col < 9; act_col++ {
		if s.matrix[r][act_col] == v {
			return false
		}
	}

	// validate the number is not present in the subgrid
	row_start := r - (r % 3)
	col_start := c - (c % 3)
	for i := row_start; i < row_start+3; i++ {
		for j := col_start; j < col_start+3; j++ {
			if s.matrix[i][j] == v {
				return false
			}
		}
	}

	return true
}

func (m Sudoku) setVal(r, c, v int) {
	m.matrix[r][c] = v
}

func generateSudoku() *Sudoku {

	n := 9
	H := make([][]int, n)
	for i := 0; i < len(H); i++ {
		H[i] = make([]int, n)
	}

	sudoku := &Sudoku{H}
	sudoku.setVal(0, 3, 4)
	sudoku.setVal(0, 6, 1)
	sudoku.setVal(0, 7, 8)
	sudoku.setVal(0, 8, 6)
	sudoku.setVal(1, 1, 8)
	sudoku.setVal(1, 2, 6)
	sudoku.setVal(1, 4, 1)
	sudoku.setVal(1, 5, 3)
	sudoku.setVal(1, 7, 7)
	sudoku.setVal(1, 8, 2)
	sudoku.setVal(2, 0, 2)
	sudoku.setVal(2, 1, 1)
	sudoku.setVal(2, 3, 8)
	sudoku.setVal(3, 0, 7)
	sudoku.setVal(3, 2, 8)
	sudoku.setVal(3, 7, 1)
	sudoku.setVal(3, 8, 3)
	sudoku.setVal(5, 3, 7)
	sudoku.setVal(5, 5, 9)
	sudoku.setVal(6, 1, 4)
	sudoku.setVal(6, 2, 1)
	sudoku.setVal(6, 5, 7)
	sudoku.setVal(6, 6, 6)
	sudoku.setVal(7, 1, 5)
	sudoku.setVal(7, 6, 8)
	sudoku.setVal(7, 8, 7)
	sudoku.setVal(8, 1, 7)
	sudoku.setVal(8, 2, 9)
	sudoku.setVal(8, 4, 3)
	sudoku.setVal(8, 6, 2)
	sudoku.setVal(8, 8, 1)

	return sudoku
}

func printMatrix(matrix [][]int) {

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {

			if (r == 3 || r == 6 || r == 9) && c == 0 {
				fmt.Printf("\n")
			}

			fmt.Printf("%v ", matrix[r][c])
			if c != len(matrix[0])-1 {
				fmt.Print("")
				if c == 2 || c == 5 {
					fmt.Print(" ")
				}
			} else {
				fmt.Printf("\n")
			}
		}
	}
	fmt.Printf("\n")
}
