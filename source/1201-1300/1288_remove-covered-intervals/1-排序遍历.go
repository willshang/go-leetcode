package main

import "sort"

func main() {

}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	maxValue := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == intervals[i-1][0] { // 合并
			res++
			continue
		}
		if intervals[i][1] > maxValue {
			maxValue = intervals[i][1] // 更新
		} else {
			res++ // 合并
		}
	}
	return len(intervals) - res
}
