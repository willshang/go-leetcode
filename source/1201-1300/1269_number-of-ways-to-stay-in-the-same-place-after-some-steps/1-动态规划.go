package main

func main() {

}

// leetcode1269_停在原地的方案数
var mod = 1000000007

func numWays(steps int, arrLen int) int {
	length := min(arrLen-1, steps)
	dp := make([][]int, steps+1) // dp[i][j] => 在i步操作之后，指针位于下标j的方案数
	for i := 0; i <= steps; i++ {
		dp[i] = make([]int, length+1)
	}
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j <= length; j++ {
			dp[i][j] = dp[i-1][j] // 不动
			if j >= 1 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod // 右移
			}
			if j <= length-1 {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod // 左移
			}
		}
	}
	return dp[steps][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
