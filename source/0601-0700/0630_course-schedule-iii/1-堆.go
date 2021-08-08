package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode630_课程表III
func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	sum := 0
	for i := 0; i < len(courses); i++ {
		count, endTime := courses[i][0], courses[i][1]
		if sum+count <= endTime { // 时间充足，学习
			sum = sum + count
			heap.Push(&intHeap, count)
		} else if intHeap.Len() > 0 && count < intHeap[0] {
			// 当前花费时间比之前时间最大耗时少，放弃之前最大耗时的课程
			top := heap.Pop(&intHeap).(int) // 最大放弃
			sum = sum - top                 // 减去最大
			sum = sum + count               // 添加当前
			heap.Push(&intHeap, count)
		}
	}
	return intHeap.Len()
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
