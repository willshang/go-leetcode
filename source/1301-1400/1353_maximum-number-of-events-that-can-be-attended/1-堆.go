package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 5},
		{1, 5},
		{1, 5}, {2, 3}, {2, 3}}
	fmt.Println(maxEvents(arr))
}

func maxEvents(events [][]int) int {
	n := len(events)
	res := 0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < n; i++ {
		heap.Push(&intHeap, events[i])
	}
	cur := 1 // 从1开始
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([]int)
		if node[1] < cur { // 当前
			continue
		}
		if node[0] >= cur || node[0] == node[1] {
			res++
			cur = node[0] + 1
			continue
		}
		node[0] = cur
		heap.Push(&intHeap, node)
	}
	return res
}

type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
