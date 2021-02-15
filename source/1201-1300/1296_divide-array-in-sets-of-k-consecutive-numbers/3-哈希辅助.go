package main

import "sort"

func main() {

}

// leetcode1296_划分数组为连续数字的集合
func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}
	if k == 1 {
		return true
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		value := m[nums[i]]
		if value > 0 {
			for j := 0; j < k; j++ {
				if m[nums[i]+j] < value {
					return false
				}
				m[nums[i]+j] = m[nums[i]+j] - value
			}
		}
	}
	return true
}
