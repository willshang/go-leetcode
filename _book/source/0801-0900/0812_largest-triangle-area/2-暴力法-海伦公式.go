package main

import (
	"fmt"
	"math"
)

func main() {
	points := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Println(largestTriangleArea(points))
}

// leetcode812_最大三角形面积
func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				// p = (a+b+c)/2
				// s = p*(p-a)*(p-b)*(p-c)
				a := length(points[i], points[j])
				b := length(points[i], points[k])
				c := length(points[j], points[k])
				p := (a + b + c) / 2
				area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

// 求两点距离
func length(p1, p2 []int) float64 {
	l := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	return math.Sqrt(float64(l))
}
