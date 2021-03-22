package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(maximumScore([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}))
}

var dp [][]int

func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp = make([][]int, m) // dp[i][j] 左i右j的值
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	res := dfs(nums, multipliers, 0, n-1)
	return res
}

func dfs(nums []int, multipliers []int, i, j int) int {
	n, m := len(nums), len(multipliers)
	if dp[i][j-n+m] != math.MinInt32 {
		return dp[i][j-n+m]
	}
	if j-i == n-m {
		dp[i][j-n+m] = max(dp[i][j-n+m], nums[i]*multipliers[m-1])
		dp[i][j-n+m] = max(dp[i][j-n+m], nums[j]*multipliers[m-1])
		return dp[i][j-n+m]
	}
	k := i + n - 1 - j
	left := dfs(nums, multipliers, i+1, j) + nums[i]*multipliers[k]
	right := dfs(nums, multipliers, i, j-1) + nums[j]*multipliers[k]
	dp[i][j-n+m] = max(left, right)
	return dp[i][j-n+m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
