package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}

// leetcode410_分割数组的最大值
func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	sub := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	// dp[i][j]表示前i个数字被分割成j段的结果
	// 0<=k<i枚举所有可以被分成j-1段的情况
	// 前 k个数被分割为j−1段，而第 k+1到第i个数为第j段
	// dp[i][j] = min{max(dp[k][j-1], sum(k+1, i))}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		// 分成m段，可能不够分，最多分min(i,m)
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
