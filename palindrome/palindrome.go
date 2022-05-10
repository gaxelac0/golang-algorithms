package main

import "fmt"

func main() {
	v := []int{0, 1, 2, 1, 0, 0}
	b := palindrome(v, 0, len(v)-1)
	fmt.Printf("%v is palindrome? %v", v, b)
}

func palindrome(v []int, start, end int) bool {

	if start >= end {
		return true
	} else if v[start] != v[end] {
		return false
	} else {
		return palindrome(v, start+1, end-1)
	}
}
