package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
}

func smallestRangeI(A []int, K int) int {
	if len(A) == 1 {
		return 0
	}
	sort.Ints(A)
	if A[len(A)-1]-A[0] > 2*K {
		return A[len(A)-1] - A[0] - 2*K
	}
	return 0
}
