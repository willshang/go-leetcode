package main

func main() {

}

// leetcode799_香槟塔
func champagneTower(poured int, query_row int, query_glass int) float64 {
	n, m := query_row, query_glass
	dp := make([][]float64, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]float64, n+2)
	}
	dp[0][0] = float64(poured) // 初始值
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][j] > 1 {
				dp[i+1][j] = dp[i+1][j] + (dp[i][j]-1)/2.0     // 往左分
				dp[i+1][j+1] = dp[i+1][j+1] + (dp[i][j]-1)/2.0 // 往右分
			}
		}
	}
	if dp[n][m] > 1 {
		return 1.0
	}
	return dp[n][m]
}
