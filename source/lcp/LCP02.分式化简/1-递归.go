package main

import "fmt"

func main() {
	fmt.Println(fraction([]int{3, 2, 0, 2}))
}

func fraction(cont []int) []int {
	if len(cont) == 1 {
		return []int{cont[0], 1}
	}
	n := fraction(cont[1:])
	m := cont[0]
	return []int{m*n[0] + n[1], n[0]}
}
