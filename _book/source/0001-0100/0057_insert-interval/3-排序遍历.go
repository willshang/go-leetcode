package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

// leetcode57_插入区间
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		arr := res[len(res)-1]
		if intervals[i][0] > arr[1] {
			res = append(res, intervals[i])
		} else if intervals[i][1] > arr[1] {
			res[len(res)-1][1] = intervals[i][1]
		}
	}
	return res
}
