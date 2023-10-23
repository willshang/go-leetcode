package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(3, 14))
}

func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		dp[1][i] = i // 1个鸡蛋N层楼，需要移动N次
	}
	for i := 1; i <= K; i++ {
		dp[i][1] = 1 // i个鸡蛋1层楼，只需要移动1次
	}
	for i := 2; i <= K; i++ {
		for j := 2; j <= N; j++ {
			if dp[i][j] == 0 {
				dp[i][j] = N // 最多N次，默认值
			}
			for x := 1; x <= j; x++ { // x是目标楼层，不断尝试x,找到最小值
				// dp[i][j-x] 没碎  dp[i-1][x-1] 碎了， 次数+1
				value := max(dp[i][j-x], dp[i-1][x-1]) + 1
				dp[i][j] = min(dp[i][j], value)
			}
		}
	}
	return dp[K][N]
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
