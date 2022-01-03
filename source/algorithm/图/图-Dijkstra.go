package main

import "math"

func main() {

}

// Dijkstra：index => 其它地址距离
// arr：邻接矩阵
func Dijkstra(arr [][]int, index int) []int {
	n := len(arr)
	maxValue := math.MaxInt32
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[index] = 0
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
	return dis
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
