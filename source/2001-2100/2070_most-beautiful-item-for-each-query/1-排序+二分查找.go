package main

import "sort"

func main() {

}

func maximumBeauty(items [][]int, queries []int) []int {
	n := len(queries)
	m := len(items)
	res := make([]int, n)
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	for i := 1; i < m; i++ {
		items[i][1] = max(items[i][1], items[i-1][1]) // 更新最大美丽值
	}
	for i := 0; i < n; i++ {
		if queries[i] >= items[m-1][0] {
			res[i] = items[m-1][1]
			continue
		}
		if queries[i] < items[0][0] {
			continue
		}
		index := binSearch(items, queries[i]) // 二分查找
		res[i] = items[index][1]
	}
	return res
}

func binSearch(arr [][]int, target int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid][0] == target {
			left = mid + 1
		} else if arr[mid][0] < target {
			left = mid + 1
		} else if arr[mid][0] > target {
			right = mid
		}
	}
	return left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
