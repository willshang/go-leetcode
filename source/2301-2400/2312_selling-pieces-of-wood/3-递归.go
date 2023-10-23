package main

func main() {

}

var temp map[[2]int]int
var dp [][]int64

func sellingWood(m int, n int, prices [][]int) int64 {
	dp = make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	temp = make(map[[2]int]int) // 转为map方便查找
	for i := 0; i < len(prices); i++ {
		a, b, c := prices[i][0], prices[i][1], prices[i][2]
		temp[[2]int{a, b}] = c // [a,b] => c
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = int64(temp[[2]int{i, j}])
			// 2个方向进行切割
			for k := 1; k < i; k++ {
				dp[i][j] = max(dp[i][j], dp[k][j]+dp[i-k][j])
			}
			for k := 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[i][j-k])
			}
		}
	}
	return dfs(m, n)
}

func dfs(m int, n int) int64 {
	if dp[m][n] != -1 {
		return dp[m][n]
	}
	res := int64(temp[[2]int{m, n}])
	// 2个方向进行切割
	for k := 1; k < m; k++ {
		res = max(res, dfs(k, n)+dfs(m-k, n))
	}
	for k := 1; k < n; k++ {
		res = max(res, dfs(m, k)+dfs(m, n-k))
	}
	dp[m][n] = res
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
