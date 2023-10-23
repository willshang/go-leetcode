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
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		l, r := intervals[i][0], intervals[i][1]
		if left <= l && r <= right { // 合并
			res++
		}
		if right < r { // 更新
			left = l
			right = r
		}
	}
	return len(intervals) - res
}
