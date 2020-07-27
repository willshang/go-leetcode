package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	return majority(nums, 0, len(nums)-1)
}

func count(nums []int, target int, start int, end int) int {
	countNum := 0
	for i := start; i <= end; i++ {
		if nums[i] == target {
			countNum++
		}
	}
	return countNum
}

func majority(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}

	mid := (start + end) / 2

	left := majority(nums, start, mid)
	right := majority(nums, mid+1, end)
	if left == right {
		return left
	}

	leftCount := count(nums, left, start, end)
	rightCount := count(nums, right, start, end)
	if leftCount > rightCount {
		return left
	}
	return right
}
