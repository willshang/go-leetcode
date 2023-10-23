package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 3},
		{0, -4},
	}
	fmt.Println(maxProductPath(arr))
}

func maxProductPath(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, m)
	}
	dp[0][0][0] = grid[0][0] // 负数
	dp[0][0][1] = grid[0][0] // 正数
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j != 0 {
				dp[i][j][0] = min(dp[i][j-1][0]*grid[i][j], dp[i][j-1][1]*grid[i][j])
				dp[i][j][1] = max(dp[i][j-1][0]*grid[i][j], dp[i][j-1][1]*grid[i][j])
			} else if i != 0 && j == 0 {
				dp[i][j][0] = min(dp[i-1][j][0]*grid[i][j], dp[i-1][j][1]*grid[i][j])
				dp[i][j][1] = max(dp[i-1][j][0]*grid[i][j], dp[i-1][j][1]*grid[i][j])
			} else if i != 0 && j != 0 {
				if grid[i][j] > 0 {
					dp[i][j][0] = min(min(dp[i-1][j][0], dp[i][j-1][0]), min(dp[i-1][j][1], dp[i][j-1][1])) * grid[i][j]
					dp[i][j][1] = max(max(dp[i-1][j][0], dp[i][j-1][0]), max(dp[i-1][j][1], dp[i][j-1][1])) * grid[i][j]
				} else {
					dp[i][j][0] = max(max(dp[i-1][j][0], dp[i][j-1][0]), max(dp[i-1][j][1], dp[i][j-1][1])) * grid[i][j]
					dp[i][j][1] = min(min(dp[i-1][j][0], dp[i][j-1][0]), min(dp[i-1][j][1], dp[i][j-1][1])) * grid[i][j]
				}
			}
		}
	}
	a, b := dp[n-1][m-1][0], dp[n-1][m-1][1]
	if a == 0 || b == 0 {
		return max(0, max(a, b))
	}
	if a > 0 && b > 0 {
		return max(a, b) % 1000000007
	}
	if b > 0 {
		return b % 1000000007
	}
	if a > 0 {
		return a % 1000000007
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
