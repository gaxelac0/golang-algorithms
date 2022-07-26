package main

import (
	"fmt"

	"github.com/gaxelac0/golang-algorithms/util"
)

func main() {

	visitados := util.MakeSet()

	excluidos := util.MakeSet()

	graph := util.MakeGraph(10)
	graph.AddBoth(0, 1)
	graph.AddBoth(0, 1)
	graph.AddBoth(0, 6)
	graph.AddBoth(0, 4)
	graph.AddBoth(1, 6)
	graph.AddBoth(1, 3)
	graph.AddBoth(1, 2)
	graph.AddBoth(2, 3)
	graph.AddBoth(2, 8)
	graph.AddBoth(2, 6)
	graph.AddBoth(3, 8)
	graph.AddBoth(4, 6)
	graph.AddBoth(4, 5)
	graph.AddBoth(5, 6)
	graph.AddBoth(5, 9)
	graph.AddBoth(6, 9)
	graph.AddBoth(6, 7)
	graph.AddBoth(6, 8)
	graph.AddBoth(7, 8)
	graph.AddBoth(7, 9)

	fmt.Printf("Visitados := %s", backtracking(graph, 0, graph.Vertices(), visitados, excluidos).String())

}

func backtracking(g *util.CustomGraph, stage int, vertices, visitados, excluidos *util.CustomSet) *util.CustomSet {

	if stage == g.Container.Order() {
		fmt.Println("reached maximum stage: " + visitados.String())
		return visitados
	}

	for !vertices.Vacio() {

		//decision
		v := vertices.Elegir()
		vertices.Sacar(v)

		if visitados.Pertenece(v) || excluidos.Pertenece(v) {
			continue
		}
		visitados.Add(v)

		lastExcluded := util.MakeSet()
		vecinos := g.Neighbors(v)
		for !vecinos.Vacio() {
			aux := vecinos.Elegir()
			vecinos.Sacar(aux)
			excluidos.Add(aux)
			lastExcluded.Add(aux)
		}

		//backtracking
		backtracking(g, stage+1, vertices, visitados, excluidos)

		// undo decision
		for !lastExcluded.Vacio() {
			x := lastExcluded.Elegir()
			lastExcluded.Sacar(x)
			excluidos.Sacar(x)
		}
		visitados.Sacar(v)
	}

	return visitados
}

func menos(cs1, cs2 *util.CustomSet) *util.CustomSet {

	result := make(map[int]struct{})

	for k, _ := range cs1.Container {
		if _, exists := cs2.Container[k]; !exists {
			result[k] = struct{}{}
		}
	}

	return &util.CustomSet{
		Container: result,
	}
}
