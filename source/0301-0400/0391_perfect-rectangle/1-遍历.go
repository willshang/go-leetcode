package main

import (
	"fmt"
	"math"
)

func main() {

}

// leetcode391_完美矩形
func isRectangleCover(rectangles [][]int) bool {
	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := math.MaxInt32, math.MinInt32
	res := 0
	m := make(map[string]bool)
	for i := 0; i < len(rectangles); i++ {
		a, b, c, d := rectangles[i][0], rectangles[i][1], rectangles[i][2], rectangles[i][3]
		minX = min(minX, a)
		minY = min(minY, b)
		maxX = max(maxX, c)
		maxY = max(maxY, d)
		res = res + (c-a)*(d-b)
		arr := []string{
			fmt.Sprintf("%d-%d", a, b),
			fmt.Sprintf("%d-%d", c, d),
			fmt.Sprintf("%d-%d", a, d),
			fmt.Sprintf("%d-%d", c, b),
		}
		for _, v := range arr {
			if _, ok := m[v]; ok {
				delete(m, v)
			} else {
				m[v] = true
			}
		}
	}
	if (maxX-minX)*(maxY-minY) != res { // 满足条件1：所有小矩形相加的面积之和要等于完美矩形的面积
		return false
	}
	if len(m) != 4 || // 满足条件2：除最终大矩形4个顶点外其它点都出现偶数次
		m[fmt.Sprintf("%d-%d", minX, minY)] == false ||
		m[fmt.Sprintf("%d-%d", minX, maxY)] == false ||
		m[fmt.Sprintf("%d-%d", maxX, minY)] == false ||
		m[fmt.Sprintf("%d-%d", maxX, maxY)] == false {
		return false
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
