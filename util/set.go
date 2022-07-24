package util

type CustomSet struct {
	Container map[int]struct{}
}

func MakeSet() *CustomSet {
	return &CustomSet{
		Container: make(map[int]struct{}),
	}
}

func (cs *CustomSet) Add(v int) {
	if !cs.Pertenece(v) {
		cs.Container[v] = struct{}{}
	}
}

func (cs *CustomSet) Sacar(v int) {
	pertenece := cs.Pertenece(v)
	if pertenece {
		delete(cs.Container, v)
	}
}

func (cs *CustomSet) Pertenece(v int) bool {
	_, pertenece := cs.Container[v]
	return pertenece
}

func (cs *CustomSet) Elegir() int {
	for k := range cs.Container {
		return k
	}
	return -1
}

func (cs *CustomSet) String() {

	for k, _ := range cs.Container {
		print(k)
	}
}

func (cs *CustomSet) Copiar() *CustomSet {
	return &CustomSet{
		Container: cs.Container,
	}
}

func (cs *CustomSet) Vacio() bool {
	return (len(cs.Container) == 0)
}
