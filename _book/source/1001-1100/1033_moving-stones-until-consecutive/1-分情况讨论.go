package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStones(1, 2, 5))
}

// leetcode1033_移动石子直到连续
func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	a, b, c = arr[0], arr[1], arr[2]
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
