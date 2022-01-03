package main

import "sort"

func main() {

}

// leetcode2122_还原原数组
func recoverArray(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[0] || (nums[i]-nums[0])%2 == 1 || nums[i] == nums[i-1] {
			continue
		}
		// leetcode954.二倍数对数组
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[nums[i]]++
		}
		res := make([]int, 0)
		k := (nums[i] - nums[0]) / 2
		for j := 0; j < n; j++ {
			if m[nums[j]] == 0 {
				continue
			}
			if m[nums[j]+2*k] == 0 {
				break
			}
			m[nums[j]]--
			m[nums[j]+2*k]--
			res = append(res, nums[j]+k)
		}
		if len(res) == n/2 {
			return res
		}
	}
	return nil
}
