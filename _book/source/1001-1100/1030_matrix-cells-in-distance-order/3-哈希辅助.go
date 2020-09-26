package main

import (
	"fmt"
)

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0)
	m := make(map[int][][]int)
	max := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			length := abs(i, r0) + abs(j, c0)
			m[length] = append(m[length], []int{i, j})
			if length > max {
				max = length
			}
		}
	}
	for i := 0; i <= max; i++ {
		for j := 0; j < len(m[i]); j++ {
			res = append(res, m[i][j])
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
