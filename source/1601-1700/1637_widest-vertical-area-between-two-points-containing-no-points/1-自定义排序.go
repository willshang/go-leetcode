package main

import "sort"

func main() {

}

// leetcode1637_两点之间不包含任何点的最宽垂直面积
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	res := 0
	for i := 1; i < len(points); i++ {
		res = max(res, points[i][0]-points[i-1][0])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
