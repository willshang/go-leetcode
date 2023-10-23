package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(int(1e10))
}

// leetcode1976_到达目的地的方案数
var mod = 1000000007

func countPaths(n int, roads [][]int) int {
	if n <= 2 {
		return 1
	}
	maxValue := int(1e12)   // math.MaxInt32=2147483647 才10位<200*1e9，可以使用11位或者更大，这里使用11不行
	arr := make([][]int, n) // 邻接矩阵：i=>j的最短距离
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = maxValue
		}
		arr[i][i] = 0
	}
	for i := 0; i < len(roads); i++ {
		a, b, c := roads[i][0], roads[i][1], roads[i][2]
		arr[a][b] = c
		arr[b][a] = c
		dp[a][b] = 1
		dp[b][a] = 1
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if arr[i][k] == maxValue { // 不联通
				continue
			}
			for j := 0; j < n; j++ {
				// 不联通/距离大于当前值
				if arr[j][k] == maxValue || arr[i][k]+arr[k][j] > arr[i][j] {
					continue
				}
				if arr[i][k]+arr[k][j] < arr[i][j] {
					dp[i][j] = (dp[i][k] * dp[k][j]) % mod
					arr[i][j] = arr[i][k] + arr[k][j]
				} else if arr[i][k]+arr[k][j] == arr[i][j] {
					dp[i][j] = (dp[i][j] + dp[i][k]*dp[k][j]) % mod
				}
			}
		}
	}
	return dp[0][n-1]
}
