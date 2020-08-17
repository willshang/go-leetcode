package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

// leetcode72_编辑距离
func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	// dp[i][j]代表 word1的i位置转换成word2的j位置需要最少步数
	dp := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
	}
	dp[0][0] = 0
	// 到word2[0]需要全部删除，有多少删除多少
	for i := 1; i <= n1; i++ {
		dp[i][0] = i
	}
	// 到word2[i]需要添加，有多少添加多少
	for i := 1; i <= n2; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 相同不需要操作
			} else {
				dp[i][j] = dp[i-1][j-1] + 1            // 替换
				dp[i][j] = min(dp[i][j], dp[i][j-1]+1) // 插入
				dp[i][j] = min(dp[i][j], dp[i-1][j]+1) // 删除
			}
		}
	}
	return dp[n1][n2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
