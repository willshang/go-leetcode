package main

import "fmt"

func main() {
	area := [][]int{{1, 2}, {3, 4}}
	fmt.Println(projectionArea(area))
}

// leetcode883_三维形体投影面积
// 1.xy面，grid[i][j]>0的个数累加
// 2.xz面, 行的最大值累加
// 3.yz面, 列的最大值累加
func projectionArea(grid [][]int) int {
	yz := [51]int{}
	xz := [51]int{}
	res := 0
	for i, line := range grid {
		for j, k := range line {
			if k == 0 {
				continue
			}
			res++
			yz[i] = max(yz[i], k)
			xz[j] = max(xz[j], k)
		}
	}
	for i := range yz {
		res = res + yz[i] + xz[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
