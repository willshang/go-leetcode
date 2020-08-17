package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}

func maxCoins(nums []int) int {
	n := len(nums)
	arr := make([]int, n+2)
	arr[0], arr[len(arr)-1] = 1, 1
	for i := 1; i <= n; i++ {
		arr[i] = nums[i-1]
	}
	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
	}
	// dp[i][j] 表示填满开区间(i,j)能得到的最多硬币数
	// i => left
	// k => i
	// j => right
	for j := 2; j <= n+1; j++ {
		for i := j - 2; i >= 0; i-- {
			for k := i + 1; k < j; k++ {
				sum := arr[i] * arr[k] * arr[j]
				sum = sum + dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
