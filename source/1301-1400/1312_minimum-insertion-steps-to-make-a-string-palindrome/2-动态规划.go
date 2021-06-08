package main

func main() {

}

func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n) // dp[i][j]表示字符串s[i:j]成为回文的添加次数
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for k := 2; k <= n; k++ { // 枚举长度
		for i := 0; i <= n-k; i++ { // 枚举起始点
			j := i + k - 1
			dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1 // 需要添加一个
			if s[i] == s[j] {
				dp[i][j] = min(dp[i][j], dp[i+1][j-1]) // 相等就不需要添加
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
