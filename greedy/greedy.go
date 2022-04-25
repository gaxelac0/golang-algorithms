package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

/**
Greedy algorithm that solves the problem of change with a specific set of coins.
Parameter: vuelto integer value expression for the change to be returned.
Parameter: coinType is an descendant ordered array of coin denominations to be used

Functions:
 - Selection: select coinType of greater denomination first.
 - Factibility: cannot exceed the _vuelto_
 - Solution: minimum quantity of coins to return the _vuelto_ as exactly as posible.

Conditions:
 - Give back the minimum quantity of coins possible.
 - Choose always the maximum denomination of coinType posible.

 Temporal Cost:
**/
func coinChangeProblem(vuelto float64, coinType []float64) []int {

	var cambio float64 = 0
	var coins = make([]int, len(coinType))
	for idx := 0; cambio < vuelto && idx < len(coinType); {
		if coinType[idx]+float64(cambio) <= float64(vuelto) {
			coins[idx] += 1
			cambio += coinType[idx]
		} else {
			idx++
		}
	}
	return coins
}

func main() {
	var change float64 = 15.5
	var coinType = []float64{10, 5, 2, 1, 0.5} // order it if needed.
	r := coinChangeProblem(change, coinType)
	fmt.Printf("Change to return: %v\n", change)
	fmt.Printf("Types: %+v\n", coinType)
	fmt.Printf("Result: %+v\n", r)
}
