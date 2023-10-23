package main

func main() {

}

// leetcode712_两个字符串的最小ASCII删除和
func minimumDeleteSum(s1 string, s2 string) int {
	a, b := len(s1), len(s2)
	// 最长公共子序列
	dp := make([][]int, a+1)
	for i := 0; i <= a; i++ {
		dp[i] = make([]int, b+1)
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return sumAscii(s1) + sumAscii(s2) - 2*dp[a][b]
}

func sumAscii(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res + int(s[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
