package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func maxTwoEvents(events [][]int) int {
	res := 0
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	prevMaxValue := 0
	for i := 0; i < len(events); i++ {
		a, b, c := events[i][0], events[i][1], events[i][2]
		for intHeap.Len() > 0 && intHeap[0][1] < a { // 之前的结束时间 小于 当前的开始时间
			value := heap.Pop(&intHeap).([3]int)[2] // 小根堆取值：按结束时间从小到大
			prevMaxValue = max(prevMaxValue, value) // 更新之前最大值
		}
		res = max(res, prevMaxValue+c)
		heap.Push(&intHeap, [3]int{a, b, c})
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type IntHeap [][3]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
