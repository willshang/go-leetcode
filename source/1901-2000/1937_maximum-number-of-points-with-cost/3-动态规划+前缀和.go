package main

import (
	"fmt"
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
	dp := make([]int, m) // dp[i][j]表示在第j列选择的格子的最大分数
	// dp[i][j]=max{dp[i−1][j']-abs(j-j')} + points[i][j]
	// j > j' => dp[i][j]=max{dp[i−1][j']+j'} + points[i][j] - j
	// j <= j' => dp[i][j]=max{dp[i−1][j']-j'} + points[i][j] + j
	for i := 0; i < n; i++ {
		temp := make([]int, m)
		leftArr := make([]int, m)  // 左边最大值
		rightArr := make([]int, m) // 右边最大值
		leftArr[0] = dp[0]
		rightArr[m-1] = dp[m-1] - (m - 1)
		for j := 1; j < m; j++ {
			leftArr[j] = max(leftArr[j-1], dp[j]+j)
		}
		for j := m - 2; j >= 0; j-- {
			rightArr[j] = max(rightArr[j+1], dp[j]-j)
		}
		for j := 0; j < m; j++ {
			temp[j] = points[i][j] + max(leftArr[j]-j, rightArr[j]+j)
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

// leetcode1937_扣分后的最大得分
func maxPoints(points [][]int) int64 {
	n, m := len(points), len(points[0])
	dp := make([]int, m) // dp[j]表示在第j列选择的格子的最大分数
	// dp[i][j]=max{dp[i−1][j']-abs(j-j')} + points[i][j]
	// j > j' => dp[i][j]=max{dp[i−1][j']+j'} + points[i][j] - j
	// j <= j' => dp[i][j]=max{dp[i−1][j']-j'} + points[i][j] + j
	for i := 0; i < n; i++ {
		temp := make([]int, m)
		leftArr := make([]int, m)  // 左边最大值
		rightArr := make([]int, m) // 右边最大值
		leftArr[0] = dp[0]
		rightArr[m-1] = dp[m-1] - (m - 1)
		for j := 1; j < m; j++ {
			leftArr[j] = max(leftArr[j-1], dp[j]+j)
		}
		for j := m - 2; j >= 0; j-- {
			rightArr[j] = max(rightArr[j+1], dp[j]-j)
		}
		for j := 0; j < m; j++ {
			temp[j] = points[i][j] + max(leftArr[j]-j, rightArr[j]+j)
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
