package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))
}

func maxResult(nums []int, k int) int {
	n := len(nums)
	if k > n {
		k = n
	}
	res := nums[0]
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [2]int{0, nums[0]})
	for i := 1; i < len(nums); i++ {
		for i-intHeap[0][0] > k { // 不满足删除
			heap.Pop(&intHeap)
		}
		res = intHeap[0][1] + nums[i]
		heap.Push(&intHeap, [2]int{i, res})
	}
	return res
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
