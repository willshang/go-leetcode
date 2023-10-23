package main

import (
	"container/heap"
)

func main() {

}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	Heap := &NodeHeap{}
	heap.Init(Heap)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(Heap, Node{
				i: nums1[i],
				j: nums2[j],
			})
			if Heap.Len() > k {
				heap.Pop(Heap)
			}
		}
	}
	res := make([][]int, 0)
	for Heap.Len() > 0 {
		node := heap.Pop(Heap).(Node)
		res = append(res, []int{node.i, node.j})
	}
	return res
}

type Node struct {
	i int
	j int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].i+h[i].j > h[j].i+h[j].j
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
