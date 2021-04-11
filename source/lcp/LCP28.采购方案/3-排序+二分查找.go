package main

import "sort"

func main() {

}

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	ln := len(nums)
	res := 0
	for i := 0; i < ln; i++ {
		target := target - nums[i]
		index := search(nums[i+1:], target)
		res = res + index
	}
	return res % 1000000007
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
