package main

import "fmt"

func main() {
	/*	arr := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}*/
	arr := [][]int{
		{1, 0},
	}
	fmt.Println(uniquePathsWithObstacles(arr))
}

// dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n < 1 {
		return 0
	}
	m := len(obstacleGrid[0])
	if m < 1 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 && j != 0 {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = dp[i][j-1]
				}
			} else if i != 0 && j == 0 {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[n-1][m-1]
}
