package main

import "fmt"

func main() {
	fmt.Println(minDays(10))
}

// leetcode1553_吃掉N个橘子的最少天数
var dp map[int]int

func minDays(n int) int {
	dp = make(map[int]int)
	dp[0] = 0
	dp[1] = 1
	return dfs(n)
}

func dfs(n int) int {
	if value, ok := dp[n]; ok {
		return value
	}
	// 吃n/2的情况，先吃掉n%2，然后就剩下dfs(n/2)+1
	// 吃n/3的情况，先吃点n%3, 然后就剩下dfs(n/3)+1
	dp[n] = min(dfs(n/2)+n%2+1, dfs(n/3)+n%3+1)
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
