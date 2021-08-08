package main

import (
	"container/heap"
	"math"
)

func main() {

}

// leetcode1675_数组的最小偏移量
func minimumDeviation(nums []int) int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	minValue := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 { // 奇数x2=>变为偶数放入堆， 统一处理为偶数
			nums[i] = nums[i] * 2
		}
		heap.Push(&intHeap, nums[i])
		minValue = min(minValue, nums[i]) // 记录最小值
	}
	res := intHeap[0] - minValue
	for intHeap.Len() > 0 && intHeap[0]%2 == 0 { //  把最大偶数处理/2
		node := heap.Pop(&intHeap).(int)
		minValue = min(minValue, node/2)
		heap.Push(&intHeap, node/2)
		res = min(res, intHeap[0]-minValue) // 目标结果：将最大值除以2，用最大值减去最小值
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
	return h[i] > h[j]
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
