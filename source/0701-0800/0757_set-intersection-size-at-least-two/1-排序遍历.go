package main

import "sort"

func main() {

}

// leetcode757_设置交集大小至少为2
func intersectionSizeTwo(intervals [][]int) int {
	n := len(intervals)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 2 // 每个区间还需要找到的交点的个数默认为2
	}
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	for i := n - 1; i >= 0; i-- {
		start := intervals[i][0]
		for k := start; k < start+arr[i]; k++ { // 一般最多取前面2个（start、start+1）
			for j := i - 1; j >= 0; j-- { // 往前遍历
				if arr[j] > 0 && k <= intervals[j][1] { // // 当前的start或者start+1小于前面的end，交点arr[j]-1
					arr[j]--
				}
			}
		}
		res = res + arr[i]
	}
	return res
}
