package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(largest1BorderedSquare(arr))
}

// leetcode1139_最大的以1为边界的正方形
func largest1BorderedSquare(grid [][]int) int {
	res := 0
	arr := [100][100][2]int{}
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 {
				arr[i][j][0] = 1
			} else {
				arr[i][j][0] = arr[i-1][j][0] + 1 // 当前坐标上边1的长度
			}
			if j == 0 {
				arr[i][j][1] = 1
			} else {
				arr[i][j][1] = arr[i][j-1][1] + 1 // 当前坐标左边1的长度
			}
			minValue := min(arr[i][j][0], arr[i][j][1]) // 左边、上边连续1选短的
			for k := minValue; k > res; k-- {
				if arr[i][j-k+1][0] >= k && arr[i-k+1][j][1] >= k { // 判断另外2条边
					res = k
					break
				}
			}
		}
	}
	return res * res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
