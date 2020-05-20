package main

import "fmt"

func main() {
	points := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Println(largestTriangleArea(points))
}

func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if area(points[i], points[j], points[k]) > maxArea {
					maxArea = area(points[i], points[j], points[k])
				}
			}
		}
	}
	return maxArea
}

// 三角形面积=|(x1 * y2 + x2 * y3 + x3 * y1 - y1 * x2 - y2 * x3 - y3 * x1)|/2
func area(p1, p2, p3 []int) float64 {
	return abs(
		p1[0]*p2[1]+p2[0]*p3[1]+p3[0]*p1[1]-
			p1[0]*p3[1]-p2[0]*p1[1]-p3[0]*p2[1]) / 2
}

func abs(num int) float64 {
	if num < 0 {
		num = -num
	}
	return float64(num)
}
