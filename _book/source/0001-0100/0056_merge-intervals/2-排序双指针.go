package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); {
		end := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
			j++
		}
		res = append(res, []int{intervals[i][0], end})
		i = j
	}
	return res
}
