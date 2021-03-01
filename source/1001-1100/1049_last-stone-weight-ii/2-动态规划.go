package main

func main() {

}

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + stones[i]
	}
	// 求最后1个最小，把石头分2堆，求差值
	// 题目转换为0-1背包问题，容量为sum/2，能装多大体积
	target := sum / 2
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j-stones[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i-1]]+stones[i-1])
			}
		}
	}
	return sum - 2*dp[n][sum/2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
