package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxPoints([][]int{
		{1, 5},
		{2, 3},
		{4, 2},
	}))
}

func maxPoints(points [][]int) int64 {
	n, m := len(points), len(points[0])
	dp := make([][]int, n) // dp[i][j]表示在第i行第j列选择的格子的最大分数
	// dp[i][j]=max{dp[i−1][j']-abs(j-j')} + points[i][j]
	// j > j' => dp[i][j]=max{dp[i−1][j']+j'} + points[i][j] - j
	// j <= j' => dp[i][j]=max{dp[i−1][j']-j'} + points[i][j] + j
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for j := 0; j < m; j++ {
		dp[0][j] = points[0][j]
	}
	for i := 1; i < n; i++ {
		maxValue := math.MinInt32 // max{dp[i−1][j']+j'}
		for j := 0; j < m; j++ {  // 正序, j越大=>maxValue越大
			maxValue = max(maxValue, dp[i-1][j]+j)
			// dp[i][j]=max{dp[i−1][j']+j'} + points[i][j] - j
			dp[i][j] = max(dp[i][j], maxValue+points[i][j]-j)
		}
		maxValue = math.MinInt32      // max{dp[i−1][j']-j'}
		for j := m - 1; j >= 0; j-- { // 逆序, j越小=>maxValue越大
			maxValue = max(maxValue, dp[i-1][j]-j)
			// dp[i][j]=max{dp[i−1][j']-j'} + points[i][j] + j
			dp[i][j] = max(dp[i][j], maxValue+points[i][j]+j)
		}
	}
	res := 0
	for j := 0; j < m; j++ {
		res = max(res, dp[n-1][j])
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
