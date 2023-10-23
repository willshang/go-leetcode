package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	arr := make([]int, len(nums)+1)
	arr[0] = 0
	for i := 1; i <= len(nums); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	for i := 0; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == k {
				res++
			}
		}
	}
	return res
}
