package main

func main() {

}

// leetcode664_奇怪的打印机
var dp [][]int

func strangePrinter(s string) int {
	n := len(s)
	dp = make([][]int, n) // dp[i][j] => 打印S[i], S[i+1], ..., S[j]所需的次数
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	return dfs(s, 0, n-1)
}

func dfs(s string, i, j int) int {
	if i > j {
		return 0
	}
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	res := dfs(s, i+1, j) + 1 // 单独打印i
	for k := i + 1; k <= j; k++ {
		if s[i] == s[k] { // 相同的时候，打印i-k
			res = min(res, dfs(s, i, k-1)+dfs(s, k+1, j))
		}
	}
	dp[i][j] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
