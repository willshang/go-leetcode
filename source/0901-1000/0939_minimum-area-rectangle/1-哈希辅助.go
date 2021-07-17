package main

import (
	"fmt"
	"math"
)

func main() {

}

// leetcode939_最小面积矩形
func minAreaRect(points [][]int) int {
	res := math.MaxInt32
	m := make(map[string]bool)
	n := len(points)
	for i := 0; i < n; i++ {
		a, b := points[i][0], points[i][1]
		m[fmt.Sprintf("%d,%d", a, b)] = true
	}
	for i := 0; i < n; i++ {
		a, b := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			c, d := points[j][0], points[j][1]
			if a == c || b == d {
				continue
			}
			area := abs((a - c) * (b - d))
			target1 := fmt.Sprintf("%d,%d", a, d) // 枚举另外2个点
			target2 := fmt.Sprintf("%d,%d", c, b)
			if area < res && m[target1] == true && m[target2] == true {
				res = area
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
