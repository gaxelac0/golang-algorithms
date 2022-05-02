package main

import "fmt"

func main() {

	var m [3][4]int

	/*for r := 0; r < 3; r++ {

		for c := 0; c < 3; c++ {
			m[r][c] = 1
		}
	}*/

	fmt.Printf("%v\n", len(m))
	fmt.Printf("%v\n", len(m[0]))

}
