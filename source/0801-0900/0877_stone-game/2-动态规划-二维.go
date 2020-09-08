package main

import "fmt"

func main() {
	fmt.Println(stoneGame([]int{5, 3, 4, 5}))
}

// leetcode877_石子游戏
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 玩家得分：自己得分-对手得分
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
