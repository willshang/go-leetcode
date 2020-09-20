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

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按照结束时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1
	end := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		node := intervals[i]
		if node[0] >= end {
			end = node[1]
			count++
		}
	}
	return len(intervals) - count
}
