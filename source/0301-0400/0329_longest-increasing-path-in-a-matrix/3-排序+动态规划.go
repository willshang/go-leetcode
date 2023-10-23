package main

import (
	"sort"
)

func main() {

}

// leetcode329_矩阵中的最长递增路径
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func longestIncreasingPath(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	temp := make([][3]int, 0)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = 1
			temp = append(temp, [3]int{i, j, matrix[i][j]})
		}
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i][2] < temp[j][2]
	})
	res := 1 // 一个数的时候，没有周围4个数，此时为1
	for i := 0; i < len(temp); i++ {
		a, b := temp[i][0], temp[i][1]
		for k := 0; k < 4; k++ {
			x, y := a+dx[k], b+dy[k]
			if 0 <= x && x < n && 0 <= y && y < m && matrix[x][y] > matrix[a][b] {
				dp[x][y] = max(dp[x][y], dp[a][b]+1) // 更新长度
				res = max(res, dp[x][y])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
