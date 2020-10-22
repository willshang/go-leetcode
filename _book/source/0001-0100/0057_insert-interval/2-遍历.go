package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{6, 9}}, []int{2, 5}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
	}
	left := i
	i = len(intervals) - 1
	for ; i >= 0 && intervals[i][0] > newInterval[1]; i-- {
	}
	right := i
	if left > right {
		return append(intervals[:left], append([][]int{newInterval}, intervals[left:]...)...)
	}
	newInterval[0] = min(newInterval[0], intervals[left][0])
	newInterval[1] = max(newInterval[1], intervals[right][1])
	return append(intervals[:left], append([][]int{newInterval}, intervals[right+1:]...)...)
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
