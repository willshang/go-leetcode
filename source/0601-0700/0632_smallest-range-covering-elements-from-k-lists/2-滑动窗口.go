package main

import "math"

func main() {

}

// leetcode632_最小区间
func smallestRange(nums [][]int) []int {
	n := len(nums)
	m := make(map[int][]int)
	minValue, maxValue := math.MaxInt32, math.MinInt32
	for i := 0; i < n; i++ {
		for j := 0; j < len(nums[i]); j++ {
			minValue = min(minValue, nums[i][j])
			maxValue = max(maxValue, nums[i][j])
			m[nums[i][j]] = append(m[nums[i][j]], i) // 存的是值对应的下标
		}
	}
	res := []int{minValue, maxValue}
	left, right := minValue, minValue // 双指针
	window := make(map[int]int)       // 滑动窗口：包含n个列的时候，更新范围
	for ; right <= maxValue; right++ {
		if len(m[right]) > 0 {
			for i := 0; i < len(m[right]); i++ {
				window[m[right][i]]++ // 添加进窗口
			}
			for len(window) == n {
				if right-left < res[1]-res[0] {
					res = []int{left, right}
				}
				for i := 0; i < len(m[left]); i++ {
					window[m[left][i]]--
					if window[m[left][i]] == 0 {
						delete(window, m[left][i])
					}
				}
				left++
			}
		}
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
