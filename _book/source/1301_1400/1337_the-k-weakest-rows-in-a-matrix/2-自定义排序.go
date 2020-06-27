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

func kWeakestRows(mat [][]int, k int) []int {
	arr := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		sum := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				sum++
			}
		}
		arr = append(arr, []int{sum, i})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, arr[i][1])
	}
	return res
}
