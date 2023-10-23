package main

import "container/heap"

func main() {

}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	if len(heights) <= 1 {
		return 0
	}
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	need := 0
	for i := 1; i < len(heights); i++ {
		need = heights[i] - heights[i-1]
		if need <= 0 {
			continue
		}
		heap.Push(intHeap, need)
		if intHeap.Len() > ladders {
			need = heap.Pop(intHeap).(int)
			if need > bricks {
				return i - 1
			}
			bricks = bricks - need
		}
	}
	return len(heights) - 1
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

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
