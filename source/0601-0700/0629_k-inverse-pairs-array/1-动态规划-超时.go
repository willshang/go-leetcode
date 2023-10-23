package main

import "fmt"

func main() {
	fmt.Println(kInversePairs(3, 1))
}

var mod = 1000000007

func kInversePairs(n int, k int) int {
	dp := make([][]int, n+1) // dp[n][k]表示1-n的排列中，包含k个逆序对
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// 前面i-1个数，i个插入位
			// 插入到最后，不增加： f(i,j) = f(i,j) + f(i-1,j)
			// 插入到倒数第2个，增加：f(i,j) = f(i,j) + f(i-1,j-1)
			// ...
			// 插入到倒数第i个，增加：f(i,j) = f(i,j) + f(i-1,j-i+1)
			// f(i,j) = f(i-1,j) + f(i-1,j-1) + ... + f(i-1,j-i+1)
			for l := max(0, j-i+1); l <= j; l++ {
				dp[i][j] = (dp[i][j] + dp[i-1][l]) % mod
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
