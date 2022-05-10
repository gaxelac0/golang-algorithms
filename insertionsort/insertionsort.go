package main

import "fmt"

func main() {

	v := []int{0, 1, 4, 3, 2}
	fmt.Printf("%v - Antes\n", v)
	insertionSort(v)
	fmt.Printf("%v - Despues\n", v)
}

func insertionSort(v []int) {

	n := len(v)
	for i := 2; i < n; i++ {
		for j := i; j > 0; j-- {
			if v[j] < v[j-1] {
				aux := v[j-1]
				v[j-1] = v[j]
				v[j] = aux
			}
		}
	}
}
