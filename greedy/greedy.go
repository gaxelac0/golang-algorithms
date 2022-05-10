package main

import (
	"fmt"
	"math"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

type candidato struct {
	idx         int
	relacionVsP float64
}

func NewCandidato(idx int, relacionVsP float64) *candidato {
	return &candidato{idx: idx, relacionVsP: relacionVsP}
}

type objeto struct {
	valor, peso float64
}

func NewObjeto(valor, peso float64) *objeto {
	return &objeto{valor: valor, peso: peso}
}

func mainCoinChange() {
	var change float64 = 0.7
	var coinType = []float64{10, 5, 2, 1, 0.5} // order it if needed.
	r := coinChangeProblem(change, coinType)
	fmt.Printf("Change to return: %v\n", change)
	fmt.Printf("Types: %+v\n", coinType)
	fmt.Printf("Result: %+v\n", r)
}

func main() {

	objetos := make([]objeto, 0)
	objetos = append(objetos, *NewObjeto(7, 10))
	objetos = append(objetos, *NewObjeto(4, 2))
	objetos = append(objetos, *NewObjeto(5, 4))
	objetos = append(objetos, *NewObjeto(4, 5))

	r := backpackProblem(10, objetos)
	fmt.Printf("Result: %+v\n", r)
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

 Temporal Cost: O(n) in which n is the coinType ? TODO check this.
**/
func coinChangeProblem(vuelto float64, coinType []float64) []int {

	var cambio float64 = 0
	var coins = make([]int, len(coinType))
	idx := 0
	for cambio < vuelto && idx < len(coinType) {
		if coinType[idx]+float64(cambio) <= float64(vuelto) {
			coins[idx] += 1
			cambio += coinType[idx]
		} else {
			idx++
		}
	}

	if idx >= len(coinType) && vuelto != cambio {

		r := make([]int, len(coinType))
		for i := 0; i < len(r); i++ {
			r[i] = -1
		}

		return r
	}

	return coins
}

/**
Temporal Cost: O(nlog n)
**/
func backpackProblem(weight float64, objetos []objeto) []float64 {

	// array of size objetos
	// array containing the % of the object i'm carrying optimizing the weight
	resultado := make([]float64, len(objetos))

	// array of size objetos
	// candidatos is an array containing the relation value/weight of each object
	candidatos := calculoCandidatos(objetos)

	// accumulated weight until now
	var accum float64 = 0
	i := 0
	for accum <= weight && i < len(candidatos) { // O(n)

		// available weight
		aval := weight - accum
		idx := candidatos[i].idx
		obj := objetos[idx]
		min := math.Min(float64(1), float64(aval/obj.peso))
		resultado[idx] = min
		accum += (obj.peso * min)
		i++

	}
	return resultado
}

// lets assume we used merge sort insted of bubble sort so O(nlog n)
func calculoCandidatos(objetos []objeto) []candidato {

	candidatos := make([]candidato, 0)
	for i := 0; i < len(objetos); i++ {
		candidatos = append(candidatos, *NewCandidato(i,
			math.Min(float64(1), float64(objetos[i].valor/objetos[i].peso))))
	}

	// Here I would use a merge_sort but using bubblesort i can work with the weights
	for i := 0; i < len(candidatos)-1; i++ {
		for j := 0; j < len(candidatos); j++ {
			if candidatos[i].relacionVsP < candidatos[i+1].relacionVsP {
				candidatos[i], candidatos[i+1] = candidatos[i+1], candidatos[i]
			}
		}
	}

	return candidatos
}

func min(i0, i1 int) int {
	return int(math.Min(float64(i0), float64(i1)))
}

func max(i1, i2 float64) float64 {
	return math.Max(float64(i1), i2)
}
