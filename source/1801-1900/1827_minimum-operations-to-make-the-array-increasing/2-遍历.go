package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{1, 1, 1}))
}

func minOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	res := 0
	target := nums[0]
	for i := 1; i < n; i++ {
		if target >= nums[i] {
			res = res + target + 1 - nums[i]
			target = target + 1
		} else {
			target = nums[i]
		}
	}
	return res
}
