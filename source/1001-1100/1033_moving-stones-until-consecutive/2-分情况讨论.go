package main

import (
	"fmt"
)

func main() {
	fmt.Println(numMovesStones(1, 2, 5))
	fmt.Println(numMovesStones(5, 2, 1))
}

func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	if a < b && b < c {
		if b-a == 1 && c-b == 1 {
			return []int{0, 0}
		} else if b-a > 2 && c-b > 2 {
			return []int{2, c - a - 2}
		} else {
			return []int{1, c - a - 2}
		}
	}
	return []int{0, 0}
}
