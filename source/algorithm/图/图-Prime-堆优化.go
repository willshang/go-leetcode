package main

import (
	"container/heap"
	"math"
)

func main() {

}

// Prime-堆优化-邻接表
func Prime(arr [][]int) int {
	res := 0
	n := len(arr)
	visited := make([]bool, n)
	target := 0
	dis := make([]int, n) // 任意选择的节点：到其它点的距离
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt32
	}
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [2]int{0, target}) // [2]int{距离，下标}
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([2]int)
		minValue, index := node[0], node[1]
		if visited[index] == true {
			continue
		}
		visited[index] = true
		res = res + minValue
		for j := 0; j < len(arr[index]); j++ {
			if visited[j] == false && dis[j] > arr[index][j] {
				dis[j] = arr[index][j]
				heap.Push(&intHeap, [2]int{arr[index][j], j})
			}
		}
	}
	return res
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
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
