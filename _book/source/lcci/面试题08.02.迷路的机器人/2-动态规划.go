package main

import "fmt"

func main() {
	//arr := [][]int{
	//	{0, 0, 0},
	//	{0, 1, 0},
	//	{0, 0, 0},
	//}
	arr := [][]int{
		{0, 1},
		{0, 0},
	}
	fmt.Println(pathWithObstacles(arr))
}

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	res := make([][]int, 0)
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[n-1][m-1] == 1 {
		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
			} else if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + 1
			} else if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + 1
			} else {
				obstacleGrid[i][j] = max(obstacleGrid[i][j-1], obstacleGrid[i-1][j]) + 1
			}
		}
	}
	total := n + m - 1
	if obstacleGrid[n-1][m-1] != total {
		return res
	}
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if obstacleGrid[i][j] == total {
			newArr := make([][]int, 0)
			newArr = append(newArr, []int{i, j})
			res = append(newArr, res...)
			total = total - 1
		}
		if i == 0 && j == 0 {
			break
		}
		if i == 0 && obstacleGrid[i][j-1] == total {
			j--
		} else if j == 0 && obstacleGrid[i-1][j] == total {
			i--
		} else if obstacleGrid[i-1][j] == total {
			i--
		} else if obstacleGrid[i][j-1] == total {
			j--
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
