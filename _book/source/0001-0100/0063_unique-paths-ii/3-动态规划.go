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

// leetcode63_不同路径II
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
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 {
				if j == 0 {
					obstacleGrid[i][j] = 1
				} else {
					obstacleGrid[i][j] += obstacleGrid[i][j-1]
				}
			} else {
				if j == 0 {
					obstacleGrid[i][j] += obstacleGrid[i-1][j]
				} else {
					obstacleGrid[i][j] += obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
				}
			}
		}
	}
	return obstacleGrid[n-1][m-1]
}
