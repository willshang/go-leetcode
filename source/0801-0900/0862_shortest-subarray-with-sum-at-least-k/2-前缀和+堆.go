package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestSubarray([]int{2, -1, 2, 3, 4, -5, 1, 2, 3}, 3))
}

var arr []int

func shortestSubarray(nums []int, k int) int {
	res := math.MaxInt32
	n := len(nums)
	arr = make([]int, n+1) // 前缀和在堆里面参与比较，使用全局变量
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i <= n; i++ {
		for intHeap.Len() > 0 && arr[i]-arr[intHeap[0]] >= k {
			res = min(res, i-intHeap[0])
			heap.Pop(&intHeap)
		}
		heap.Push(&intHeap, i)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return arr[h[i]] < arr[h[j]]
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
