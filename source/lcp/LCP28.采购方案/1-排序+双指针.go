package main

import "sort"

func main() {

}

// leetcode_lcp28_采购方案
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	j := len(nums) - 1
	res := 0
	for i := 0; i < len(nums); i++ {
		for i < j {
			if nums[i]+nums[j] <= target {
				break
			}
			j--
		}
		if i < j {
			res = res + (j - i)
		}
	}
	return res % 1000000007
}
