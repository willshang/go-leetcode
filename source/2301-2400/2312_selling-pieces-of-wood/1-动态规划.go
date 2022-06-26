package main

func main() {

}

// leetcode2312_卖木头块
func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int64, n+1)
	}
	temp := make(map[[2]int]int) // 转为map方便查找
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
	return dp[m][n]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
