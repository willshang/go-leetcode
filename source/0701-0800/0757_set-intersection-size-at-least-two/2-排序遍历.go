package main

import "sort"

func main() {

}

func intersectionSizeTwo(intervals [][]int) int {
	n := len(intervals)
	arr := []int{-1, -1}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	for i := 0; i < n; i++ {
		start, end := intervals[i][0], intervals[i][1]
		a, b := arr[len(arr)-2], arr[len(arr)-1]
		if start <= a { // 上一个开始值已经包括当前start
			continue
		}
		if b < start { // 当前开始大于之前结束，把[end-1,end]包括
			arr = append(arr, end-1)
		}
		arr = append(arr, end)
	}
	return len(arr) - 2
}
