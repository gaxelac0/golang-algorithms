package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {

	wanted := []int{1, 2, 4, 7, 7, 8, 8, 11, 17, 24}

	arr := []int{7, 4, 11, 24, 17, 8, 1, 2, 7, 8}
	MergeSort(arr, 0, 9)

	fmt := fmt.Sprintf("sorted %v, wanted %v", arr, wanted)
	t.Logf(fmt)

	if !reflect.DeepEqual(arr, wanted) {
		t.Fatalf(fmt)
	}
}
