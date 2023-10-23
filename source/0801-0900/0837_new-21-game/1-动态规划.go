package main

func main() {

}

func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W) // 为当前手中牌面为i点时获胜的概率
	for i := K; i <= N && i < K+W; i++ {
		dp[i] = 1.0
	}
	/*for i := K-1; i >= 0; i--{
		for j := 1; j <= W; j++{ // 每次选择1~W
			dp[i] = dp[i] + dp[i+j]/float64(W)
		}
	}*/
	dp[K-1] = 1.0 * (float64(min(N-K+1, W))) / float64(W)
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+W+1]-dp[i+1])/float64(W)
	}
	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
