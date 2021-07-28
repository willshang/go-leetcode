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
	dp := make([]int, m) // dp[j]表示在第j列选择的格子的最大分数
	// dp[i][j]=max{dp[i−1][j']-abs(j-j')} + points[i][j]
	// j > j' => dp[i][j]=max{dp[i−1][j']+j'} + points[i][j] - j
	// j <= j' => dp[i][j]=max{dp[i−1][j']-j'} + points[i][j] + j
	for i := 0; i < n; i++ {
		temp := make([]int, m)
		maxValue := math.MinInt32 // max{dp[i−1][j']+j'}
		for j := 0; j < m; j++ {  // 正序, j越大=>maxValue越大
			maxValue = max(maxValue, dp[j]+j)
			// dp[i][j]=max{dp[i−1][j']+j'} + points[i][j] - j
			temp[j] = max(temp[j], maxValue+points[i][j]-j)
		}
		maxValue = math.MinInt32      // max{dp[i−1][j']-j'}
		for j := m - 1; j >= 0; j-- { // 逆序, j越小=>maxValue越大
			maxValue = max(maxValue, dp[j]-j)
			// dp[i][j]=max{dp[i−1][j']-j'} + points[i][j] + j
			temp[j] = max(temp[j], maxValue+points[i][j]+j)
		}
		copy(dp, temp)
	}
	res := 0
	for j := 0; j < m; j++ {
		res = max(res, dp[j])
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
