package main

import "testing"

func Test_integerSquareRoot(t *testing.T) {
	type args struct {
		start  int
		end    int
		filing int
	}
	tests := []struct {
		name   string
		args   args
		result int
	}{
		{
			name: "EntireSquareRoot of 4 / sqrt(4)",
			args: args{
				start:  1,
				end:    2,
				filing: 4,
			},
			result: 2,
		},
		{
			name: "EntireSquareRoot of 16 / sqrt(16)",
			args: args{
				start:  1,
				end:    4,
				filing: 16,
			},
			result: 4,
		},
		{
			name: "EntireSquareRoot of 1 / sqrt(1)",
			args: args{
				start:  1,
				end:    1,
				filing: 1,
			},
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerSquareRoot(tt.args.start, tt.args.end, tt.args.filing); got != tt.result {
				t.Errorf("integerSquareRoot() = %v, result %v", got, tt.result)
			}
		})
	}
}
