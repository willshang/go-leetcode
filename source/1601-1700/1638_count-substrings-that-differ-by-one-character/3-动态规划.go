package main

func main() {

}

func countSubstrings(s, t string) int {
	res := 0
	m, n := len(s), len(t)
	// dp以s[i]和t[j]结尾的所有子串对中，满足恰好只有一个字符不同的字符串对的数目
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 以s[i]和t[j]为结尾的子串，最多有多少个连续相同的字符
	same := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		same[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i+1][j+1] + dp[i][j]
				same[i+1][j+1] = same[i][j] + 1
			} else {
				dp[i+1][j+1] = same[i][j] + 1
			}
			res = res + dp[i+1][j+1]
		}
	}
	return res
}
