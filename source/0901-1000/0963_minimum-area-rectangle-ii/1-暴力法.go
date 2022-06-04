package main

import "math"

func main() {

}

func minAreaFreeRect(points [][]int) float64 {
	res := math.MaxFloat64
	n := len(points)
	m := map[[2]int]bool{}
	for i := 0; i < n; i++ {
		m[[2]int{points[i][0], points[i][1]}] = true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				for k := j + 1; k < n; k++ {
					if i != k {
						x1, x2, x3 := points[i][0], points[j][0], points[k][0]
						y1, y2, y3 := points[i][1], points[j][1], points[k][1]
						// 计算x,y坐标 矩形1243
						x4 := x2 + x3 - x1
						y4 := y2 + y3 - y1
						if m[[2]int{x4, y4}] == true {
							area := (x2 - x1) * (x3)
						}
					}
				}
			}
		}
	}
	return res
}
