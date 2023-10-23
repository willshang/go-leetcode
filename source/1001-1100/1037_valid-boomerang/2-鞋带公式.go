package main

import "fmt"

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
}

// 鞋带公式
// S=|(x1 * y2 + x2 * y3 + x3 * y1 - y1 * x2 - y2 * x3 - y3 * x1)|/2
// S!=0组成三角形
func isBoomerang(points [][]int) bool {
	return points[0][0]*points[1][1]+points[1][0]*points[2][1]+points[2][0]*points[0][1] !=
		points[0][1]*points[1][0]+points[1][1]*points[2][0]+points[2][1]*points[0][0]
}
