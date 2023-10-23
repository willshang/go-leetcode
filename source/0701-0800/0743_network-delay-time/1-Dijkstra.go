package main

import "math"

func main() {

}

func networkDelayTime(times [][]int, n int, k int) int {
	maxValue := math.MaxInt32 / 10
	arr := make([][]int, n) // 邻接矩阵：i=>j的最短距离
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = maxValue
		}
	}
	for i := 0; i < len(times); i++ {
		a, b, c := times[i][0]-1, times[i][1]-1, times[i][2] // a=>b
		arr[a][b] = c
	}
	dis := make([]int, n) // k到其他点的距离
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[k-1] = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		target := -1 // 寻找未访问的距离起点最近点
		for j := 0; j < len(visited); j++ {
			if visited[j] == false && (target == -1 || dis[j] < dis[target]) {
				target = j
			}
		}
		visited[target] = true
		for j := 0; j < len(arr[target]); j++ { // 更新距离
			dis[j] = min(dis[j], dis[target]+arr[target][j])
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if dis[i] == maxValue {
			return -1
		}
		res = max(res, dis[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
