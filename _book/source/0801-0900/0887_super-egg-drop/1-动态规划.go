package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(3, 14))
}

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}
	for j := 1; j <= N; j++ {
		for i := 1; i <= K; i++ {
			// dp[i-1][j-1](碎了)+dp[i][j-1](没碎)
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= N {
				return j
			}
		}
	}
	return N
}
