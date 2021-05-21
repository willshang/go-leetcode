package main

func main() {

}

var mod = 1000000007

func numWays(steps int, arrLen int) int {
	length := min(arrLen-1, steps)
	dp := make([]int, length+1) // dp[j] => 指针位于下标j的方案数
	dp[0] = 1
	for i := 1; i <= steps; i++ {
		temp := make([]int, length+1)
		for j := 0; j <= length; j++ {
			temp[j] = dp[j] // 不动
			if j >= 1 {
				temp[j] = (temp[j] + dp[j-1]) % mod // 右移
			}
			if j <= length-1 {
				temp[j] = (temp[j] + dp[j+1]) % mod // 左移
			}
		}
		dp = temp
	}
	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
