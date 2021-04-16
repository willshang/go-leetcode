package main

import "fmt"

func main() {
	fmt.Println(numOfArrays(2, 3, 1))
}

// leetcode1420_生成数组
var mod = 1000000007

func numOfArrays(n int, m int, k int) int {
	if k == 0 {
		return 0
	}
	dp := make([][][]int, n+1) // 数组第i位最大值为j，比较次数为k的结果
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}
	for i := 1; i <= m; i++ {
		dp[1][i][1] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for a := 1; a <= k; a++ {
				for b := 1; b < j; b++ { // 比j小,才可以次数+1
					dp[i][j][a] = (dp[i][j][a] + dp[i-1][b][a-1]) % mod
				}
				dp[i][j][a] = (dp[i][j][a] + dp[i-1][j][a]*j) % mod // 跟j相同，可选择[1,j]共j个数
			}
		}
	}
	res := 0
	for i := 1; i <= m; i++ {
		res = (res + dp[n][i][k]) % mod
	}
	return res
}
