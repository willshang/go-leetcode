package main

import "container/heap"

func main() {

}

// leetcode1845_座位预约管理系统
type SeatManager struct {
	intHeap IntHeap
}

func Constructor(n int) SeatManager {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 1; i <= n; i++ {
		heap.Push(&intHeap, i)
	}
	return SeatManager{intHeap: intHeap}
}

func (this *SeatManager) Reserve() int {
	top := heap.Pop(&this.intHeap).(int)
	return top
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.intHeap, seatNumber)
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
