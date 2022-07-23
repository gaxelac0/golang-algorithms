package main

import "testing"

func TestSudoku_solve(t *testing.T) {
	type fields struct {
		matrix [][]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sudoku{
				matrix: tt.fields.matrix,
			}
			if got := s.solve(); got != tt.want {
				t.Errorf("Sudoku.solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
