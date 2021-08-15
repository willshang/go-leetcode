package main

import "container/heap"

func main() {

}

// leetcode1962_移除石子使总数最小
func minStoneSum(piles []int, k int) int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(piles); i++ {
		heap.Push(&intHeap, piles[i])
	}
	for i := 1; i <= k; i++ {
		node := heap.Pop(&intHeap).(int)
		value := node - node/2
		heap.Push(&intHeap, value)
	}
	res := 0
	for i := 0; i < len(piles); i++ {
		res = res + intHeap[i]
	}
	return res
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
