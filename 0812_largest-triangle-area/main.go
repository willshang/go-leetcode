package main

import "fmt"

func main() {
	points := [][]int{{0,0},{0,1},{1,0},{0,2},{2,0}}
	fmt.Println(largestTriangleArea(points))
}
func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	n := len(points)
	for i := 0; i < n; i++{
		for j := i+1; j < n; j++{
			for k := j+1; k < n; k++{
				maxArea = max(maxArea,area(points[i],points[j],points[k]))
			}
		}
	}
	return maxArea
}


// S=(1/2)*(x1y2+x2y3+x3y1-x1y3-x2y1-x3y2)
func area(p1,p2,p3 []int) float64  {
	return abs(
		p1[0]*p2[1]+p2[0]*p3[1]+p3[0]*p1[1] -
			p1[0]*p3[1]-p2[0]*p1[1]-p3[0]*p2[1])/2
}

func abs(num int)float64  {
	if num < 0{
		num = -num
	}
	return float64(num)
}

func max(a,b float64) float64  {
	if a > b {
		return a
	}
	return b
}