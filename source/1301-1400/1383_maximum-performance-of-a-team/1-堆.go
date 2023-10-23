package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode1383_最大的团队表现值
func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	arr := make([]Node, 0)
	for i := 0; i < len(speed); i++ {
		arr = append(arr, Node{
			speed:      speed[i],
			efficiency: efficiency[i],
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].efficiency == arr[j].efficiency {
			return arr[i].speed > arr[j].speed
		}
		return arr[i].efficiency > arr[j].efficiency // 效率递减
	})
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	res := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		s := arr[i].speed
		e := arr[i].efficiency
		sum = sum + s          // 速度和
		heap.Push(&intHeap, s) // 速度
		if intHeap.Len() > k {
			value := heap.Pop(&intHeap).(int)
			sum = sum - value
		}
		res = max(res, sum*e)
	}
	return res % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	speed      int
	efficiency int
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
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
