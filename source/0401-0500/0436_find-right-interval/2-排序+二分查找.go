package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRightInterval([][]int{{1, 2}}))
}

func findRightInterval(intervals [][]int) []int {
	m := make(map[int]int)
	n := len(intervals)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
		m[intervals[i][0]] = i // 存储start对应的下标，因为起点都不相同
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < n; i++ {
		target := intervals[i][1]
		index := m[intervals[i][0]]
		left, right := i, n-1
		for left <= right {
			mid := left + (right-left)/2
			if target <= intervals[mid][0] {
				res[index] = m[intervals[mid][0]]
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return res
}
