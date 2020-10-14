package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 3},
		{-2, 2},
	}
	fmt.Println(kClosest(arr, 1))
}

func kClosest(points [][]int, K int) [][]int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(points); i++ {
		heap.Push(&intHeap, points[i])
	}
	res := make([][]int, 0)
	for i := 0; i < K; i++ {
		value := heap.Pop(&intHeap).([]int)
		res = append(res, value)
	}
	return res
}

type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] <
		h[j][0]*h[j][0]+h[j][1]*h[j][1]
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
