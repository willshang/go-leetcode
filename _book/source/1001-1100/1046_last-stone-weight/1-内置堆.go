package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func lastStoneWeight(stones []int) int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(stones); i++ {
		heap.Push(&intHeap, stones[i])
	}
	for intHeap.Len() > 1 {
		a := heap.Pop(&intHeap).(int)
		b := heap.Pop(&intHeap).(int)
		if a > b {
			heap.Push(&intHeap, a-b)
		}
	}
	if intHeap.Len() > 0 {
		res := heap.Pop(&intHeap).(int)
		return res
	}
	return 0
}
