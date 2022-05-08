package main

import (
	"fmt"
)

func main() {

	// conn := 0
	// for i := 999999999; i > 0; i-- {
	// 	conn = conn + i
	// }

	// fmt.Printf("%v", conn)

	slarkHunting()
}

func printMatrix() {

	var m [3][4]int

	/*for r := 0; r < 3; r++ {

		for c := 0; c < 3; c++ {
			m[r][c] = 1
		}
	}*/

	fmt.Printf("%v\n", len(m))
	fmt.Printf("%v\n", len(m[0]))

}
