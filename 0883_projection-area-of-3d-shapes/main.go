package main

import "fmt"

func main() {
	area := [][]int{{1, 2}, {3, 4}}
	fmt.Println(projectionArea(area))
}
func projectionArea(grid [][]int) int {
	xs := [51]int{}
	ys := [51]int{}
	res := 0

	for i, line := range grid {
		for j, k := range line {
			if k == 0 {
				continue
			}
			res++
			xs[i] = max(xs[i], k)
			ys[j] = max(ys[j], k)
		}
	}

	for i := range xs {
		res = res + xs[i] + ys[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
