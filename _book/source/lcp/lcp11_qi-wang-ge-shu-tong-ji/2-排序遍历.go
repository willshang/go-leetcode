package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(expectNumber([]int{1, 2, 3}))
}

func expectNumber(scores []int) int {
	sort.Ints(scores)
	count := 0
	for i := 1; i < len(scores); i++ {
		if scores[i] == scores[i-1] {
			count++
		}
	}
	return len(scores) - count
}
