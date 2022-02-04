package main

func main() {

}

// leetcode2140_解决智力问题
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1) // dp[i] => 解决第i道题后的最高分数
	for i := 0; i < n; i++ {
		next := min(n, i+questions[i][1]+1)
		dp[i+1] = max(dp[i+1], dp[i])                   // 跳过
		dp[next] = max(dp[next], dp[i]+questions[i][0]) // 不跳过
	}
	return int64(dp[n])
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
