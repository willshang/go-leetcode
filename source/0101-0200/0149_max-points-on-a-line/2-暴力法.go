package main

import "fmt"

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
}

// leetcode149_直线上最多的点数
func maxPoints(points [][]int) int {
	res := 0
	n := len(points)
	if n < 3 {
		return n
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			count := 2
			x1 := points[i][0] - points[j][0]
			y1 := points[i][1] - points[j][1]
			for k := j + 1; k < n; k++ {
				x2 := points[i][0] - points[k][0]
				y2 := points[i][1] - points[k][1]
				if x1*y2 == x2*y1 { // 斜率相同+1
					count++
				}
			}
			if count > res {
				res = count
			}
		}
	}
	return res
}
