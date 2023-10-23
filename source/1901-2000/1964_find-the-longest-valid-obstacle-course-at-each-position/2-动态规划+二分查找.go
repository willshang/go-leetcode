package main

import "sort"

func main() {

}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	res := make([]int, n)
	dp := make([]int, 0)
	for i := 0; i < n; i++ {
		index := sort.SearchInts(dp, obstacles[i]+1)
		if index < len(dp) {
			dp[index] = obstacles[i]
		} else {
			dp = append(dp, obstacles[i])
		}
		res[i] = index + 1
	}
	return res
}
