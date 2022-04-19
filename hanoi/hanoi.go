package main

import (
	"fmt"
)

/**
Hanoi is a (traditional game) function that orders a series of disks of different diameters in different pins.
We have A, B and C pins, we have the disks stored in the A disk in decreasing-size order from bottom to top.

The disks are moved from A pin to C pin using B as an auxiliar.
 ===================================
|        _     (1)		|		|	|
|		___    (n-2) 	|       |	|
|	   _____   (n-1)    |       |	|
|	  _______  (n)		|       |	|
|		 A       	    B       C   |
 ===================================
Rules:
	- One disk at the time
	- Only the top of each stack can be moved
	- Disk must be ordered by diameter. (n) > (n-1) > 1
**/
func hanoi(n int, origin *stack, aux *stack, destination *stack) {

	origin_length := len(origin.s)
	if origin_length != 0 {

		if n == 1 {
			destination.Push(origin.Pop())
		} else {
			hanoi(n-1, origin, destination, aux)
			destination.Push(origin.Pop())
			hanoi(n-1, aux, origin, destination)
		}
	}
}

func main() {
	A := NewStack()
	A.Push(3)
	A.Push(2)
	A.Push(1)
	B := NewStack()
	C := NewStack()

	fmt.Printf("Before\n")
	fmt.Printf("A = %+v\n", A.s)
	fmt.Printf("B = %+v\n", B.s)
	fmt.Printf("C = %+v\n", C.s)
	hanoi(3, A, B, C)
	fmt.Printf("After\n")
	fmt.Printf("A = %+v\n", A.s)
	fmt.Printf("B = %+v\n", B.s)
	fmt.Printf("C = %+v\n", C.s)
}

type stack struct {
	s []int
}

func NewStack() *stack {
	return &stack{make([]int, 0)}
}

func (s *stack) Push(v int) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() int {
	l := len(s.s)
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}
