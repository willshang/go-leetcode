package main

import (
	"math"
)

func main() {

}

// leetcode1584_连接所有点的最小费用
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	arr := make([][]int, n) // 邻接矩阵
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			c := dis(points[i], points[j])
			arr[i][j] = c
			arr[j][i] = c
		}
	}
	return Prime(arr)
}

// Prime：传入邻接矩阵
func Prime(arr [][]int) int {
	res := 0
	n := len(arr)
	visited := make([]bool, n)
	target := 0
	visited[target] = true // 任选1个即可
	dis := make([]int, n)  // 任意选择的节点：到其它点的距离
	for i := 0; i < n; i++ {
		dis[i] = arr[target][i]
	}
	for i := 1; i < n; i++ { // 执行n-1次：n-1条边
		var index int
		minValue := math.MaxInt32
		for j := 0; j < n; j++ { // 寻找：未访问过的最短边
			if visited[j] == false && dis[j] < minValue {
				minValue = dis[j]
				index = j
			}
		}
		visited[index] = true    // 标记为访问过的点
		res = res + minValue     // 加上最短边
		for j := 0; j < n; j++ { // 更新距离：以index为起点，更新生成树到每一个非树顶点的距离
			if visited[j] == false && dis[j] > arr[index][j] {
				dis[j] = arr[index][j]
			}
		}
	}
	return res
}

func dis(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
