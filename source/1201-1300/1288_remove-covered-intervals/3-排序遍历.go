package main

import "sort"

func main() {

}

// leetcode1288_删除被覆盖区间
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	maxValue := 0
	for i := 0; i < len(intervals); i++ {
		if maxValue < intervals[i][1] {
			res++
			maxValue = intervals[i][1]
		}
	}
	return res
}
