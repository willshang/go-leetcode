package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 1, 7}))
}

// leetcode945_使数组唯一的最小增量
func minIncrementForUnique(A []int) int {
	res := 0
	sort.Ints(A)
	for i := 0; i < len(A)-1; i++ {
		if A[i] >= A[i+1] {
			value := A[i] - A[i+1] + 1
			res = res + value
			A[i+1] = A[i+1] + value
		}
	}
	return res
}
