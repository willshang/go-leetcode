package main

import "fmt"

func main() {
	fmt.Println(stoneGameV([]int{6, 2, 3, 4, 5, 5}))
}

// leetcode1563_石子游戏V
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	sum := make([]int, n+1)
	sum[0] = stoneValue[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + stoneValue[i]
	}
	dp := make([][]int, n+1) // dp[i][j]：代表从i到j区间,Alice分数最大值
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}
	for length := 2; length <= n; length++ { // 长度从2开始到N依次枚举
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			for k := i; k <= j; k++ {
				if i > k || k+1 > j {
					continue
				}
				left := dp[i][k]
				right := dp[k+1][j]
				leftSum := sum[k]
				if i > 0 {
					leftSum = sum[k] - sum[i-1]
				}
				rightSum := sum[j] - sum[k]
				if leftSum == rightSum {
					dp[i][j] = max(dp[i][j], max(left, right)+leftSum)
				} else if leftSum > rightSum {
					dp[i][j] = max(dp[i][j], right+rightSum)
				} else if leftSum < rightSum {
					dp[i][j] = max(dp[i][j], left+leftSum)
				}
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
