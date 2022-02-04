package main

func main() {

}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1) // dp[i] => 解决第i道题后的最高分数
	// 反向
	for i := n - 1; i >= 0; i-- {
		next := min(n, i+questions[i][1]+1)
		dp[i] = max(dp[i+1], dp[next]+questions[i][0])
	}
	return int64(dp[0])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
