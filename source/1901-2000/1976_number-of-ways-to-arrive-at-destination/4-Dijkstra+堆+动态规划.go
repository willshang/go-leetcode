package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := [][]int{
		{0, 6, 7},
		{0, 1, 2},
		{1, 2, 3},
		{1, 3, 3},
		{6, 3, 3},
		{3, 5, 1},
		{6, 5, 1},
		{2, 5, 1},
		{0, 4, 5},
		{4, 6, 2},
	}
	fmt.Println(countPaths(7, arr))
}

var mod = 1000000007

func countPaths(n int, roads [][]int) int {
	mathValue := int(1e12)
	arr := make([][][2]int, n)
	for i := 0; i < len(roads); i++ { // 邻接表
		a, b, c := roads[i][0], roads[i][1], roads[i][2]
		arr[a] = append(arr[a], [2]int{b, c})
		arr[b] = append(arr[b], [2]int{a, c})
	}
	dp := make([]int, n)
	dis := make([]int, n) // 0到其他点的距离
	for i := 0; i < n; i++ {
		dis[i] = mathValue
	}
	dis[0] = 0
	dp[0] = 1
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [2]int{0, 0})

	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([2]int)
		a, d := node[0], node[1]
		if dis[a] < d {
			continue
		}
		for i := 0; i < len(arr[a]); i++ {
			b, c := arr[a][i][0], arr[a][i][1]
			if dis[a]+c < dis[b] {
				dis[b] = dis[a] + c
				dp[b] = dp[a]
				heap.Push(&intHeap, [2]int{b, dis[b]})
			} else if dis[a]+c == dis[b] {
				dp[b] = (dp[b] + dp[a]) % mod
			}
		}
	}
	return dp[n-1] % mod
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
