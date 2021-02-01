package main

import "sort"

func main() {

}

// leetcode452_用最少数量的箭引爆气球
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	right := points[0][1]
	res := 1
	for i := 0; i < len(points); i++ {
		if points[i][0] > right {
			right = points[i][1]
			res++
		}
	}
	return res
}
