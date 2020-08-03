package main

import "fmt"

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 || t < 0 {
		return false
	}
	m := make(map[int]int)
	width := t + 1
	for i := 0; i < len(nums); i++ {
		key := getKey(nums[i], width)
		if _, ok := m[key]; ok {
			return true
		}
		if value, ok := m[key-1]; ok && abs(nums[i], value) < width {
			return true
		}
		if value, ok := m[key+1]; ok && abs(nums[i], value) < width {
			return true
		}
		m[key] = nums[i]
		if i >= k {
			// 满足i和j的差的绝对值也小于等于ķ
			delete(m, getKey(nums[i-k], width))
		}
	}
	return false
}

func getKey(value, width int) int {
	if value < 0 {
		return (value+1)/width - 1
	}
	return value / width
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
