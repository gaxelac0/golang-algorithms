package main

import (
	"testing"
)

func TestMonoDecreasing(t *testing.T) {

}

func Test_findRoot(t *testing.T) {
	type args struct {
		f     a_function
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"mock", args{f: mock, start: 1024, end: 2048}, 1637},
		{"mock2", args{f: mock2, start: 64, end: 128}, 115},
		{"mock3", args{f: mock3, start: 2048, end: 4096}, 2048}, // worst case O(n)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRoot(tt.args.f, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("findRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
