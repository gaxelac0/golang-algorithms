package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	wanted := []int{2, 11, 22, 32, 36, 64}

	arr := []int{32, 64, 22, 11, 2, 36}
	QuickSort(arr, 0, len(arr)-1)

	format := fmt.Sprintf("sorted %v, wanted %v", arr, wanted)
	t.Log(format)

	if !reflect.DeepEqual(arr, wanted) {
		t.Fatalf(format)
	}
}
