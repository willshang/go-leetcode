package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(int(1e10))
}

var mod = 1000000007
var dp []int

func countPaths(n int, roads [][]int) int {
	maxValue := int(1e12)   // math.MaxInt32=2147483647 才10位<200*1e9，可以使用11位或者更大，这里使用11不行
	arr := make([][]int, n) // 邻接矩阵：i=>j的最短距离
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = maxValue
		}
		arr[i][i] = 0
	}
	for i := 0; i < len(roads); i++ {
		a, b, c := roads[i][0], roads[i][1], roads[i][2]
		arr[a][b] = c
		arr[b][a] = c
	}
	arr = Floyd(arr)
	path := make([][]int, n)
	dp = make([]int, n)
	for i := 0; i < n; i++ {
		path[i] = make([]int, 0)
		dp[i] = -1
	}
	// 构建图，都是从下标0开始
	for i := 0; i < len(roads); i++ {
		a, b, c := roads[i][0], roads[i][1], roads[i][2]
		if arr[0][b]-arr[0][a] == c {
			path[a] = append(path[a], b)
		} else if arr[0][a]-arr[0][b] == c {
			path[b] = append(path[b], a)
		}
	}
	return dfs(path, 0)
}

func dfs(path [][]int, index int) int {
	if index == len(path)-1 {
		return 1
	}
	if dp[index] != -1 {
		return dp[index]
	}
	dp[index] = 0
	for i := 0; i < len(path[index]); i++ {
		next := path[index][i]
		dp[index] = (dp[index] + dfs(path, next)) % mod
	}
	return dp[index]
}

func Floyd(arr [][]int) [][]int {
	n := len(arr)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i][k]+arr[k][j] < arr[i][j] {
					arr[i][j] = arr[i][k] + arr[k][j]
				}
			}
		}
	}
	return arr
}
