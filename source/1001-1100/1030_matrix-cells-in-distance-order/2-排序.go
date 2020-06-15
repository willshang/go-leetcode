package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
}

// leetcode1030_距离顺序排列矩阵单元格
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return length(res[i][0], r0)+length(res[i][1], c0) <
			length(res[j][0], r0)+length(res[j][1], c0)
	})
	return res
}

func length(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
