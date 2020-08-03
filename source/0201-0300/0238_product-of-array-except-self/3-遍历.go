package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

// leetcode238_除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	value := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * value
		value = value * nums[i]
	}
	return res
}
