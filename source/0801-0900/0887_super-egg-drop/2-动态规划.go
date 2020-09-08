package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(3, 14))
}

func superEggDrop(K int, N int) int {
	// dp[i][j]: N层楼， K个鸡蛋
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, K+1)
		dp[i][1] = i // i层楼1个鸡蛋，需要i次
	}
	for i := 1; i <= N; i++ {

		for i := 1; i <= K; i++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= N {
				return j
			}
		}
	}
	return N
}
