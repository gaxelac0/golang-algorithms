package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubleSort(t *testing.T) {

	list := []int{7, 12, 111, 3}
	sorted := bubbleSort(list)

	wanted := []int{3, 7, 12, 111}

	format := fmt.Sprintf("sorted %v, wanted %v", sorted, wanted)
	t.Log(format)

	if !reflect.DeepEqual(sorted, wanted) {
		t.Fatalf(format)
	}
}
