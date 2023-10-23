package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}))
}

// leetcode435_无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按照结束时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	dp := make([]int, len(intervals))
	dp[0] = 1
	res := 1
	for i := 1; i < len(intervals); i++ {
		count := 0
		for j := i - 1; j >= 0; j-- {
			if intervals[j][1] <= intervals[i][0] {
				count = max(dp[j], count)
			}
		}
		dp[i] = count + 1
		res = max(res, dp[i])
	}
	return len(intervals) - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
