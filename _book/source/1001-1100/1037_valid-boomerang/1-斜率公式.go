package main

import "fmt"

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
}

// leetcode1037_有效的回旋镖
// k1=(y1-y0)/(x1-x0) = k2 = (y2-y1)/(x2-x1)
// (x1-x0)*(y2-y1) = (x2-x1)*(y1-y0)
func isBoomerang(points [][]int) bool {
	return (points[1][0]-points[0][0])*(points[2][1]-points[1][1]) !=
		(points[2][0]-points[1][0])*(points[1][1]-points[0][1])
}
