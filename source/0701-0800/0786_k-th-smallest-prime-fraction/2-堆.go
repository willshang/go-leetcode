package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for j := 1; j < n; j++ {
		heap.Push(&intHeap, []int{arr[0], arr[j], 0, j}) // 0/j 递减：分子最小，分母依次增大
	}
	for i := 1; i <= k-1; i++ { // 取k-1个数(k从1开始)
		node := heap.Pop(&intHeap).([]int)
		x, y := node[2], node[3]
		if x+1 < y { // 下标 x+1 < y
			heap.Push(&intHeap, []int{arr[x+1], arr[y], x + 1, y})
		}
	}
	return []int{intHeap[0][0], intHeap[0][1]}
}

type IntHeap [][]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0]*h[j][1] < h[i][1]*h[j][0] } // a*d < b*c
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
