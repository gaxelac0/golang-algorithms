package main

import "fmt"

/**
pseudoM it's an algorithm that search for candidates to majoritary element in the array sent by parameter.
This is done recursively dividing the array until getting a subarray of two or one element, depending of the array's element quantity, odd or even.
Temporal Cost: O(n)

Divide & Conquer
a = 2
b = 2
k = 0

a > b^k
1 > 1

O( n ^(logb a) )
O(n ^ (log2 2))
O(n)
**/
func pseudoM(arr []int) (hasCandidate bool, lenArr int, qttyCandidate int, candidate int) {

	n := len(arr)

	if n == 1 {
		return true, 1, 1, arr[0]
	} else if n == 2 && arr[0] == arr[1] {
		return true, 2, 2, arr[0]
	} else if n > 2 {
		mid := (n - 1) / 2

		hasXI, _, cxI, xI := pseudoM(arr[0 : mid+1])
		hasXD, _, cxD, xD := pseudoM(arr[mid+1 : n])

		if hasXI {
			fmt.Printf("Izquierda %v - Candidato(%v) aparece %v veces\n", arr[0:mid+1], xI, cxI)
		} else {
			fmt.Printf("Izquierda %v - No hay candidato.\n", arr[0:mid+1])
		}

		if hasXD {
			fmt.Printf("Derecha   %v - Candidato(%v) aparece %v veces\n", arr[mid+1:n], xD, cxD)
		} else {
			fmt.Printf("Derecha   %v - No hay candidato.\n", arr[mid+1:n])
		}

		if hasXI && hasXD {

			// son el mismo elemento en los dos subarrays
			if xI == xD { // [1, 1] == [1,1] or [1] == [1] or [1] == [1, 1]
				return true, n, cxI + cxD, xI

				// no son el mismo elemento pero aparecen la misma cantidad de veces en los dos subvectores
			} else if cxD == cxI { // [1, 1] == [2,2]
				return false, n, 0, 0
			} else if cxI >= (n/2)+1 {
				return hasXI, n, cxI, xI
			} else {
				return hasXD, n, cxD, xD
			}
		} else if hasXI {
			return hasXI, n, cxI, xI
		} else if hasXD {
			return hasXD, n, cxD, xD
		} else {
			return false, n, 0, 0
		}

	} else {
		return false, n, 0, 0
	}
}

/**
Implementation of majoritary element.
A majoritary element is an element that appears at least n/2 + 1 times in the vector.
Temporal Cost: O(n)
**/
func pseudoMMajorityElement(arr []int, elem int) bool {

	cant := 0
	hasX, n, _, x := pseudoM(arr) // O(n)

	for i := 0; i < len(arr); i++ { // O(n)
		if arr[i] == x {
			cant++
		}
	}
	return hasX && elem == x && cant >= (n/2)+1
}

func main() {
	arr := []int{1, 1, 2, 2, 305, 305, 305, 305, 305}

	b := pseudoMMajorityElement(arr, 305)
	fmt.Printf("pseudoM - Mayoritario: %v\n", b)

	c := linearMajorityElement(arr, 305)
	fmt.Printf("linear - Mayoritario: %v\n", c)
}

// Precond. Array ordenado con números positivos mayores que 0.
// Costo Temporal: O(n²)
// pseudoM tries to improve that temporal cost, so it's O(n)
func linearMajorityElement(arr []int, elem int) bool {

	n := len(arr)
	var majoritaryCant int = -1
	var majoritary int
	for k := 0; k < n; k++ {

		cantK := 0
		for i := 0; i < n; i++ {
			if arr[k] == arr[i] {
				cantK++
			}
		}

		if cantK > majoritaryCant {
			majoritary = arr[k]
			majoritaryCant = cantK
		}
	}

	return majoritary == elem && (majoritaryCant >= (n/2)+1)
}
