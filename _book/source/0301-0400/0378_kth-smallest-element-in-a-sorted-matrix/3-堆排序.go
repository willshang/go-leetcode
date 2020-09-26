package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(arr, 8))
}

func kthSmallest(matrix [][]int, k int) int {
	var h IntHeap
	heap.Init(&h)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			heap.Push(&h, matrix[i][j])
		}
	}
	for h.Len() > 0 && k > 0 {
		k--
		res := heap.Pop(&h).(int)
		if k == 0 {
			return res
		}
	}
	return 0
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
