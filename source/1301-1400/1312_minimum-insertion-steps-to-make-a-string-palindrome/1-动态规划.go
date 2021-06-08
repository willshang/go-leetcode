package main

func main() {

}

func minInsertions(s string) int {
	n := len(s)
	t := reverse(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	// 求s和它的反转字符串t的 最长公共子序列的长度
	// leetcode 1143.最长公共子序列
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return n - dp[n][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(str string) string {
	arr := []byte(str)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}
