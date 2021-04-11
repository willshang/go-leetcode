package main

import "sort"

func main() {

}

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]+nums[right] > target {
			right--
		}
		res = res + right - left
		left++
	}
	return res % 1000000007
}
