package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	/*	arr := [][]int{
		{1, 0},
	}*/
	fmt.Println(uniquePathsWithObstacles(arr))
}

// dp[j] = dp[j] + dp[j-1]
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
	dp := make([]int, m)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j >= 1 && obstacleGrid[i][j-1] == 0 {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[m-1]
}
