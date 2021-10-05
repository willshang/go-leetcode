package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	dp1 := make([]int, n) // 从第一家开始打劫，最后一家不可选
	dp2 := make([]int, n) // 从第二家开始打劫，最后一家可以选
	dp1[0] = nums[0]
	dp1[1] = max(nums[0], nums[1])
	dp2[0] = 0
	dp2[1] = nums[1]
	for i := 2; i < n; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}
	return max(dp1[n-2], dp2[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
