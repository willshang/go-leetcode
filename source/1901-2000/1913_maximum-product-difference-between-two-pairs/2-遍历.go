package main

import ()

func main() {
}

func maxProductDifference(nums []int) int {
	n := len(nums)
	minA, minB := min(nums[0], nums[1]), max(nums[0], nums[1])
	maxA, maxB := max(nums[0], nums[1]), min(nums[0], nums[1])
	for i := 2; i < n; i++ {
		if nums[i] > maxA {
			maxA, maxB = nums[i], maxA
		} else if nums[i] > maxB {
			maxB = nums[i]
		}
		if nums[i] < minA {
			minA, minB = nums[i], minA
		} else if nums[i] < minB {
			minB = nums[i]
		}
	}
	return maxA*maxB - minA*minB
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
