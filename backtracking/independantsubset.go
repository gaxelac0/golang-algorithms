package main

import (
	"github.com/gaxelac0/golang-algorithms/util"
)

func main() {

	// set1 := util.MakeSet()

	// set2 := util.MakeSet()
	// set2.Add(2)

	// print(set1.Vacio())
	// print(set2.Vacio())

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

}

func backtracking(g *util.CustomGraph, stage int) *util.CustomSet {
	return nil
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
