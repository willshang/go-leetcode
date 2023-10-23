package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		if i == 0 && i+1 < n && nums[i] > nums[i+1] {
			return i
		}
		if i == n-1 && i-1 >= 0 && nums[i] > nums[i-1] {
			return i
		}
		if i-1 >= 0 && i+1 < n && nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i
		}
	}
	return -1
}
