package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(3, 14))
}

// leetcode887_鸡蛋掉落
func superEggDrop(K int, N int) int {
	// dp[i][j] 有i次操作，j个鸡蛋时能测出的最高的楼层数
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}
	for j := 1; j <= N; j++ { // 操作次数
		for i := 1; i <= K; i++ { //  K个蛋
			// dp[i][j-1](没碎)+dp[i-1][j-1](碎了)+当前
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= N {
				return j
			}
		}
	}
	return N
}
