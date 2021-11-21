package main

import "sort"

func main() {

}

// leetcode2070_每一个查询的最大美丽值
func maximumBeauty(items [][]int, queries []int) []int {
	n := len(queries)
	m := len(items)
	res := make([]int, n)
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		arr[i] = [2]int{i, queries[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	j := 0
	maxValue := 0
	for i := 0; i < n; i++ {
		index, target := arr[i][0], arr[i][1]
		for ; j < m && items[j][0] <= target; j++ {
			maxValue = max(maxValue, items[j][1])
		}
		res[index] = maxValue
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
