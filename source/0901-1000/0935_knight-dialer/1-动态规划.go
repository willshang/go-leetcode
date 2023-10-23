package main

func main() {

}

// leetcode935_骑士拨号器
func knightDialer(n int) int {
	dp := make([][10]int, n)
	for i := 0; i < 10; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][4] + dp[i-1][6]
		dp[i][1] = dp[i-1][6] + dp[i-1][8]
		dp[i][2] = dp[i-1][7] + dp[i-1][9]
		dp[i][3] = dp[i-1][4] + dp[i-1][8]
		dp[i][4] = dp[i-1][0] + dp[i-1][3] + dp[i-1][9]
		dp[i][5] = 0
		dp[i][6] = dp[i-1][0] + dp[i-1][1] + dp[i-1][7]
		dp[i][7] = dp[i-1][2] + dp[i-1][6]
		dp[i][8] = dp[i-1][1] + dp[i-1][3]
		dp[i][9] = dp[i-1][2] + dp[i-1][4]
		for j := 0; j < 10; j++ {
			dp[i][j] = dp[i][j] % 1000000007
		}
	}
	res := 0
	for i := 0; i < 10; i++ {
		res = (res + dp[n-1][i]) % 1000000007
	}
	return res
}
