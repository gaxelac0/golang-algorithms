package util

import "github.com/yourbasic/graph"

type CustomGraph struct {
	Container *graph.Mutable
}

func MakeGraph(cantNodes int) *CustomGraph {
	return &CustomGraph{
		Container: graph.New(cantNodes),
	}
}

func (cs *CustomGraph) AddBoth(v, w int) {
	cs.Container.AddBoth(v, w)
}

func (cs *CustomGraph) Vertices() *CustomSet {

	n := cs.Container.Order()

	vertices := MakeSet()
	for i := 0; i < n; i++ {
		cs.Container.Visit(i, func(w int, c int64) (skip bool) {
			vertices.Add(w)
			return
		})
	}
	return vertices
}

func (cs *CustomGraph) Neighbors(v int) *CustomSet {

	neighbors := MakeSet()
	cs.Container.Visit(v, func(w int, c int64) (skip bool) {
		neighbors.Add(w)
		return
	})

	return neighbors
}

func (cs *CustomGraph) String() string {
	return cs.Container.String()
}
