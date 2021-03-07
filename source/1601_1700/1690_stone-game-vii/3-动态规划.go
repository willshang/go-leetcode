package main

import "fmt"

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2}))
}

// leetcode1690_石子游戏VII
func stoneGameVII(stones []int) int {
	n := len(stones)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + stones[i]
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			left := arr[j+1] - arr[i+1] - dp[i+1][j] // 左边得分-得分
			right := arr[j] - arr[i] - dp[i][j-1]    // 右边得分-得分
			dp[i][j] = max(left, right)
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
