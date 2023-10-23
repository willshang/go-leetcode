package main

import "sort"

func main() {

}

// leetcode2054_两个最好的不重叠活动
func maxTwoEvents(events [][]int) int {
	arr := make([][3]int, 0)
	for i := 0; i < len(events); i++ {
		a, b, c := events[i][0], events[i][1], events[i][2]
		arr = append(arr, [3]int{a, 0, c})
		arr = append(arr, [3]int{b, 1, c}) // 拆分为2块
	}
	sort.Slice(arr, func(i, j int) bool { // 按时间排序
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})
	res := 0
	prevMaxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i][1] == 0 { // 为开始时间
			res = max(res, arr[i][2]+prevMaxValue)
		} else {
			prevMaxValue = max(prevMaxValue, arr[i][2]) // 更新之前最大值
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
