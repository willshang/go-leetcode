package main

import "container/heap"

func main() {

}

// leetcode2233_K次增加后的最大乘积
var mod = 1000000007

func maximumProduct(nums []int, k int) int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(nums); i++ {
		heap.Push(&intHeap, nums[i])
	}
	// x y：如果x>y,y+1后的结果较大
	// (x+1)*y = xy+y
	// x*(y+1) = xy+x
	for i := 1; i <= k; i++ {
		node := heap.Pop(&intHeap).(int)
		node++
		heap.Push(&intHeap, node)
	}
	res := 1
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).(int)
		res = res * node % mod
	}
	return res
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
