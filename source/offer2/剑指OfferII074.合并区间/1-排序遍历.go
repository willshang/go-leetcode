package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

// 剑指OfferII074.合并区间
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return nil
	}
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
