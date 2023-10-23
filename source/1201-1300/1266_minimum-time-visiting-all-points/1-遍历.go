package main

import "fmt"

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
}

// leetcode1266_访问所有点的最小时间
func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i++ {
		x := length(points[i][0], points[i-1][0])
		y := length(points[i][1], points[i-1][1])
		if x > y {
			res = res + x
		} else {
			res = res + y
		}
	}
	return res
}

func length(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
