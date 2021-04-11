package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(magicTower([]int{-200, -300, 400, 0}))
}

func magicTower(nums []int) int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	blood := 0
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if nums[i] < 0 {
			heap.Push(&intHeap, nums[i])
			if blood+nums[i] < 0 {
				res++
				minValue := heap.Pop(&intHeap).(int)
				blood = blood - minValue
			}
		}
		blood = blood + nums[i]
	}
	if sum < 0 {
		return -1
	}
	return res
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
