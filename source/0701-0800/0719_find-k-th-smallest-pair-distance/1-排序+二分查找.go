package main

import "sort"

func main() {

}

// leetcode719_找出第k小的距离对
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := left + (right-left)/2
		if judge(nums, mid, k) == true {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(nums []int, mid, k int) bool {
	count := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > mid {
			left++
		}
		count = count + right - left
	}
	return k <= count
}
