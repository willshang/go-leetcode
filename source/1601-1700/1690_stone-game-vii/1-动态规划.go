package main

import "fmt"

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2}))
}

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
	for j := 2; j <= n; j++ {
		for i := 0; i+j <= n; i++ {
			left := arr[i+j] - arr[i+1] - dp[i+1][i+j-1] // 左边得分-得分
			right := arr[i+j-1] - arr[i] - dp[i][i+j-2]  // 右边得分-得分
			dp[i][i+j-1] = max(left, right)
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
