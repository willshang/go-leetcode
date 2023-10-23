package main

import (
	"math"
)

func main() {

}

// leetcode743_网络延迟时间
func networkDelayTime(times [][]int, n int, k int) int {
	maxValue := math.MaxInt32 / 10
	arr := make([][][2]int, n) // 邻接表：i=>j的集合
	for i := 0; i < len(times); i++ {
		a, b, c := times[i][0]-1, times[i][1]-1, times[i][2] // a=>b
		arr[a] = append(arr[a], [2]int{b, c})
	}
	dis := make([]int, n) // k到其他点的距离
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[k-1] = 0
	queue := make([]int, 0)
	queue = append(queue, k-1)
	for len(queue) > 0 {
		a := queue[0]
		queue = queue[1:]
		for i := 0; i < len(arr[a]); i++ {
			b, c := arr[a][i][0], arr[a][i][1]
			if dis[a]+c < dis[b] { // 更新距离
				dis[b] = dis[a] + c
				queue = append(queue, b)
			}
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
