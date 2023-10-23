package main

import (
	"container/heap"
	"math"
)

func main() {

}

func networkBecomesIdle(edges [][]int, patience []int) int {
	maxValue := math.MaxInt32 / 10
	n := len(patience)
	arr := make([][][2]int, n) // 邻接表：i=>j的集合
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1] // a=>b
		arr[a] = append(arr[a], [2]int{b, 1})
		arr[b] = append(arr[b], [2]int{a, 1})
	}
	dis := make([]int, n) // k到其他点的距离
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[0] = 0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [2]int{0, 0})
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
	res := 0
	for i := 1; i < n; i++ {
		total := (2*dis[i]-1)/patience[i]*patience[i] + 2*dis[i]
		res = max(res, total)
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
