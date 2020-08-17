package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(arr, 3))
}

// leetcode378_有序矩阵中第K小的元素
func kthSmallest(matrix [][]int, k int) int {
	res := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			res = append(res, matrix[i][j])
		}
	}
	sort.Ints(res)
	return res[k-1]
}
