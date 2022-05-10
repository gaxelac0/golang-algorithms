package main

import "testing"

func Test_bbPeaks(t *testing.T) {
	type args struct {
		V     []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test-A",
			args: args{
				V:     []int{0, 1, 3, 2, 5, 1, 0},
				start: 0,
				end:   6,
			},
			want: 5,
		},
		{
			name: "Test-B",
			args: args{
				V:     []int{1, 3, 4, 5, 9, 8, 5, 3, 2, -3},
				start: 0,
				end:   9,
			},
			want: 9,
		},
		{
			name: "Test-C",
			args: args{
				V:     []int{14, 0, 1},
				start: 0,
				end:   2,
			},
			want: 14,
		},
		{
			name: "Test-D",
			args: args{
				V:     []int{14, 0, 22},
				start: 0,
				end:   2,
			},
			want: 14,
		},
		{
			name: "Test-E",
			args: args{
				V:     []int{0, 0, 22},
				start: 0,
				end:   2,
			},
			want: 0,
		},
		{
			name: "Test-F",
			args: args{
				V:     []int{-3, 0, 22},
				start: 0,
				end:   2,
			},
			want: 22,
		},
		{
			name: "Test-G",
			args: args{
				V:     []int{-3, 11, 4},
				start: 0,
				end:   2,
			},
			want: 11,
		},
		{
			name: "Test-H",
			args: args{
				V:     []int{1, 1, 1},
				start: 0,
				end:   2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := bbPeaks(tt.args.V, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("bbPeaks() = %v, want %v", got, tt.want)
			} else {
				t.Logf("\n\n========== PeakElement: %v\n\n", got)
			}
		})
	}
}
