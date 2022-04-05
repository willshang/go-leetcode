package main

import "container/heap"

func main() {

}

// leetcode2208_将数组和减半的最少操作次数
func halveArray(nums []int) int {
	n := len(nums)
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	sum := float64(0)
	for i := 0; i < n; i++ {
		sum = sum + float64(nums[i])
		heap.Push(&intHeap, []float64{float64(nums[i])})
	}
	target := sum / 2
	res := 0
	for sum > target {
		node := heap.Pop(&intHeap).([]float64)
		v := node[0]
		v = v / 2
		sum = sum - v
		heap.Push(&intHeap, []float64{v})
		res++
	}
	return res
}

type IntHeap [][]float64

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]float64)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
