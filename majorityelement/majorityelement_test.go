package main

import (
	"fmt"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		arr  []int
		elem int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{arr: []int{6, 6, 1}, elem: 6}, true},
		{"2", args{arr: []int{1, 1, 2, 3, 3, 3, 3}, elem: 3}, true},
		{"3", args{arr: []int{1, 1, 2, 2, 3, 3}, elem: 3}, false},
		{"4", args{arr: []int{7, 4, 3, 2, 1, 7, 7, 4, 7, 7, 7}, elem: 7}, true},
		{"5", args{arr: []int{1, 2, 1, 2, 1, 2, 1, 2, 2, 2, 2, 2, 1, 2}, elem: 2}, true},
		{"6 - No elements in array", args{arr: []int{}, elem: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("\nCase %v\n", tt.name)
			if got := pseudoMMajorityElement(tt.args.arr, tt.args.elem); got != tt.want {
				t.Errorf("majoritaryElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
