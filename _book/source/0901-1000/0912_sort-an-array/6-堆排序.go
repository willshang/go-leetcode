package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{5, 2, 3, 1}
	fmt.Println(sortArray(arr))
	fmt.Println(arr)
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

func sortArray(nums []int) []int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(nums); i++ {
		heap.Push(&intHeap, nums[i])
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		value := heap.Pop(&intHeap).(int)
		res = append(res, value)
	}
	return res
}
