package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var h IntHeap
	heap.Init(&h)
	for k, v := range m {
		heap.Push(&h, [2]int{k, v})
	}
	res := make([]int, 0)
	for h.Len() > 0 && k > 0 {
		k--
		node := heap.Pop(&h).([2]int)
		res = append(res, node[0])
	}
	return res
}

type IntHeap [][2]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][1] > h[j][1] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
