package main

import "sort"

func main() {

}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			res++
			left++
			right--
		} else if sum > k {
			right--
		} else {
			left++
		}
	}
	return res
}
