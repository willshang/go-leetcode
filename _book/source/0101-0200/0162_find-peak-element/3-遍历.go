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
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return n - 1
}
