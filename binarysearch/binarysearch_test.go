package main

import (
	"fmt"
	"testing"
)

func TestBs(t *testing.T) {

	elem := 5
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	bPresent := Bs(arr, elem)

	fmt := fmt.Sprintf("[without index] %d is present in arr? %v", elem, bPresent)
	if !bPresent {
		t.Fatalf(fmt)
	}
	t.Logf(fmt)
}

func TestBsIdx(t *testing.T) {

	elem := 5
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	bPresent := BsIdx(arr, 0, len(arr)-1, elem)

	fmt := fmt.Sprintf("[with index] %d is present in arr? %v", elem, bPresent)
	if !bPresent {
		t.Fatalf(fmt)
	}
	t.Logf(fmt)
}
