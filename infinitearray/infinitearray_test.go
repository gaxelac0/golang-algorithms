package main

import (
	"math"
	"testing"
)

func Test_findPos(t *testing.T) {
	type args struct {
		arr   []float64
		start int
		end   int
		x     float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1", args{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
			math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1)},
			2, 4, 5}, 4},
		{"test-2", args{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9,
			math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1)},
			8, 16, 11}, -1},
		{"test-3", args{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 11,
			math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1)},
			8, 16, 11}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Searching for: %v\n", tt.args.x)
			t.Logf("arr := %+v\n", tt.args.arr)
			t.Logf("Searching bettween %v and %v\n", tt.args.start, tt.args.end)
			if got := findPos(tt.args.arr, tt.args.start, tt.args.end, tt.args.x); got != tt.want {
				t.Errorf("findPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
