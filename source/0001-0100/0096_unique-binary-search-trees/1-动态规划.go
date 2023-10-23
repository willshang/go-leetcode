package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

// leetcode96_不同的二叉搜索树
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}
