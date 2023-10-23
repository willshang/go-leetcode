package main

import "fmt"

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
}

// leetcode523_连续的子数组和
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if k != 0 {
			sum = sum % k
		}
		if _, ok := m[sum]; ok {
			// 确保数组大小至少为2
			if i-m[sum] >= 2 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}
