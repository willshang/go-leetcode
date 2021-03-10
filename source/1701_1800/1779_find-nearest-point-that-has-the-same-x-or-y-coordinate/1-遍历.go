package main

import "math"

func main() {

}

// leetcode1779_找到最近的有相同X或Y坐标的点
func nearestValidPoint(x int, y int, points [][]int) int {
	res := -1
	maxValue := math.MaxInt32
	for i := 0; i < len(points); i++ {
		a, b := points[i][0], points[i][1]
		if a == x || b == y {
			value := abs(a-x) + abs(b-y)
			if value < maxValue {
				res = i
				maxValue = value
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
