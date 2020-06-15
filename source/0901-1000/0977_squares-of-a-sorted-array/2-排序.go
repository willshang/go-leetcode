package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

// leetcode977_有序数组的平方
func sortedSquares(A []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(A); i++ {
		res = append(res, A[i]*A[i])
	}
	sort.Ints(res)
	return res
}
