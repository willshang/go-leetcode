package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(kWeakestRows(arr, 2))
}

// leetcode1337_方阵中战斗力最弱的K行
func kWeakestRows(mat [][]int, k int) []int {
	arr := make([]int, 0)
	for i := 0; i < len(mat); i++ {
		sum := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				sum++
			}
		}
		arr = append(arr, sum*100+i)
	}
	sort.Ints(arr)
	for i := 0; i < k; i++ {
		arr[i] = arr[i] % 100
	}
	return arr[:k]
}
