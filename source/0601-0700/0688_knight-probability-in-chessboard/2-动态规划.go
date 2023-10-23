package main

import "fmt"

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
}

var dx = []int{2, 2, 1, 1, -1, -1, -2, -2}
var dy = []int{1, -1, 2, -2, 2, -2, 1, -1}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, n)
	}
	dp[row][column] = float64(1)
	for a := 1; a <= k; a++ {
		temp := make([][]float64, n)
		for i := 0; i < n; i++ {
			temp[i] = make([]float64, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for b := 0; b < 8; b++ {
					x := i + dx[b]
					y := j + dy[b]
					if 0 <= x && x < n && 0 <= y && y < n {
						temp[i][j] = temp[i][j] + dp[x][y]/8.0
					}
				}
			}
		}
		dp = temp
	}
	res := float64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res = res + dp[i][j]
		}
	}
	return res
}
