package main

func main() {

}

// leetcode837_新21点
func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W) // 为当前手中牌面为i点时获胜的概率
	var sum float64
	for i := K; i <= K+W-1; i++ {
		if i <= N {
			dp[i] = 1
		}
		sum = sum + dp[i]
	}

	for i := K - 1; i >= 0; i-- {
		dp[i] = sum / float64(W)
		sum = sum - dp[i+W] + dp[i]
	}
	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
