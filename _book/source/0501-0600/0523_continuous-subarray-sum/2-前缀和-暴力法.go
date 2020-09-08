package main

import "fmt"

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
}

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr[i] = arr[i-1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := arr[j] - arr[i] + nums[i]
			if sum == k || (k != 0 && sum%k == 0) {
				return true
			}
		}
	}
	return false
}
