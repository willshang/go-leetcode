package main

import "sort"

func main() {

}

func minimalKSum(nums []int, k int) int64 {
	sum := int64(0)
	sort.Ints(nums)
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == false && nums[i] <= k {
			sum = sum + int64(nums[i]) // 存在小于K的，加起来
			k++                        // k+1
		}
		m[nums[i]] = true
	}
	return int64((k+1)*k/2) - sum
}
