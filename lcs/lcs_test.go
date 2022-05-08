package main

import (
	"reflect"
	"testing"
)

func Test_lcs(t *testing.T) {
	type args struct {
		x []string
		y []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{

		{
			name: "1",
			args: args{
				x: []string{"", "A", "B", "A", "C", "A", "D", "A", "C"},
				y: []string{"", "A", "B", "C", "B", "D", "A"},
			},
			want: []string{"A", "B", "C", "D", "A"},
		},
		{
			name: "2",
			args: args{
				x: []string{"", "A", "B", "A", "C", "A", "D", "A", "C"},
				y: []string{"", "A", "B", "C", "B", "D", "A"},
			},
			want: []string{"A", "B", "C", "D", "A"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcs(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lcs() = %v, want %v", got, tt.want)
			}
		})
	}
}
