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

func countPaths(n int, roads [][]int) int {
	maxValue := int(1e11)   // math.MaxInt32=2147483647 才10位<200*1e9，可以使用11位或者更大
	arr := make([][]int, n) // 邻接矩阵：i=>j的最短距离
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = maxValue
		}
	}
	for i := 0; i < len(roads); i++ {
		a, b, c := roads[i][0], roads[i][1], roads[i][2]
		arr[a][b] = c
		arr[b][a] = c
	}
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[0] = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		target := -1 // 寻找未访问的距离起点最近点
		for j := 0; j < n; j++ {
			if visited[j] == false && (target == -1 || dis[j] < dis[target]) {
				target = j
			}
		}
		visited[target] = true
		for j := 0; j < n; j++ { // 更新距离
			dis[j] = min(dis[j], dis[target]+arr[target][j])
		}
	}
	// 计算某条边是否在边上
	edge := make([]int, n) // 入度
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dis[i]+arr[i][j] == dis[j] {
				edge[j]++ // 入度+1
			}
		}
	}
	dp := make([]int, n) // dp[i] 表示0=>i的最短路个数
	dp[0] = 1
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for j := 0; j < n; j++ {
			if dis[node]+arr[node][j] == dis[j] {
				dp[j] = (dp[j] + dp[node]) % mod
				edge[j]--         // 入度-1
				if edge[j] == 0 { // 入队
					queue = append(queue, j)
				}
			}
		}
	}
	return dp[n-1] % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
