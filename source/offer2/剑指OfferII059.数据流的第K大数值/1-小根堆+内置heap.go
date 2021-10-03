package main

import (
	"container/heap"
	"fmt"
)

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	kth := Constructor(k, arr)
	kth.Add(3)
	kth.Add(6)
	fmt.Println(kth.k, kth.heap)
}

// 剑指OfferII059.数据流的第K大数值
type KthLargest struct {
	k    int
	heap intHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	heap.Init(&h)
	for len(h) > k {
		heap.Pop(&h)
	}
	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (k *KthLargest) Add(val int) int {
	heap.Push(&k.heap, val)
	if len(k.heap) > k.k {
		heap.Pop(&k.heap)
	}
	return k.heap[0]
}

// 内置heap，实现接口
/*
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}
*/
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
