package main

import "fmt"

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
}

// leetcode992_K个不同整数的子数组
func subarraysWithKDistinct(nums []int, k int) int {
	return getK(nums, k) - getK(nums, k-1) // 恰好K个不同=最多K个不同-最多K-1个不同
}

// 求最多K个不同
func getK(nums []int, k int) int {
	res := 0
	n := len(nums)
	m := make(map[int]int)
	left, right := 0, 0
	count := 0
	for ; right < n; right++ {
		cur := nums[right]
		if m[cur] == 0 {
			count++ // 次数+1
		}
		m[cur]++
		for count > k {
			m[nums[left]]--
			if m[nums[left]] == 0 {
				count--
			}
			left++
		}
		res = res + right - left
	}
	return res
}
