package main

import (
	"container/heap"
	"math"
)

func main() {

}

// Dijkstra-堆优化：index => 其它地址距离
// arr => 邻接表 arr[a] = append(arr[a], [2]int{b, c}) 其中a=>b c
func Dijkstra(arr [][][2]int, index int) []int {
	n := len(arr)
	maxValue := math.MaxInt32
	dis := make([]int, n) // k到其他点的距离
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[index] = 0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [2]int{index, 0}) // 下标+距离
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([2]int) // 距离起点最近的点
		a := node[0]
		if dis[a] < node[1] { // 大于最短距离，跳过
			continue
		}
		for i := 0; i < len(arr[a]); i++ {
			b, c := arr[a][i][0], arr[a][i][1]
			if dis[a]+c < dis[b] { // 更新距离
				dis[b] = dis[a] + c
				heap.Push(&intHeap, [2]int{b, dis[b]})
			}
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
