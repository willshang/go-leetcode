package main

func main() {

}

// leetcode583_两个字符串的删除操作
func minDistance(word1 string, word2 string) int {
	a, b := len(word1), len(word2)
	// 最长公共子序列
	dp := make([][]int, a+1)
	for i := 0; i <= a; i++ {
		dp[i] = make([]int, b+1)
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return a + b - 2*dp[a][b]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
