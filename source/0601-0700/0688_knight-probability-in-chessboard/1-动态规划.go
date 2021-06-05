package main

import "fmt"

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
}

// leetcode688_“马”在棋盘上的概率
var dx = []int{2, 2, 1, 1, -1, -1, -2, -2}
var dy = []int{1, -1, 2, -2, 2, -2, 1, -1}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, n) // dp[i][j][k]在位置[i,j]移动k步
	for i := 0; i < n; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, k+1)
		}
	}
	dp[row][column][0] = float64(1)
	for a := 1; a <= k; a++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for b := 0; b < 8; b++ {
					x := i + dx[b]
					y := j + dy[b]
					if 0 <= x && x < n && 0 <= y && y < n {
						dp[i][j][a] = dp[i][j][a] + dp[x][y][a-1]/8.0
					}
				}
			}
		}
	}
	res := float64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res = res + dp[i][j][k]
		}
	}
	return res
}
