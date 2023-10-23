package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 100, 3},
		{7, 8, 9},
	}
	fmt.Println(maxValueOfCoins(arr, 2))
}

// leetcode2218_从栈中取出K个硬币的最大面值和
func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	dp := make([][]int, n+1) // 背包问题：dp[i][j]=>表示i个栈取j个硬币的最大值
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		length := len(piles[i-1])
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j]
			sum := 0
			for x := 1; x <= min(j, length); x++ { // 枚举第i个栈能取到的长度
				sum = sum + piles[i-1][x-1]                // 第i个栈的前缀和
				dp[i][j] = max(dp[i][j], dp[i-1][j-x]+sum) // x个在第i个栈取+j-x在前i-1个栈取:j减去x + x个前缀和
			}
		}
	}
	return dp[n][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
