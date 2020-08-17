package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	res = append(res, newInterval)
	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
