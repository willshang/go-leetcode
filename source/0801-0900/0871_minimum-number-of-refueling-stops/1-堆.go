package main

import "container/heap"

func main() {

}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	res := 0
	total := startFuel
	if total >= target {
		return 0
	}
	Heap := &IntHeap{}
	heap.Init(Heap)
	for i := 0; i < len(stations); i++ {
		for total < stations[i][0] {
			if Heap.Len() == 0 {
				return -1
			}
			total += heap.Pop(Heap).(int)
			res++
		}
		heap.Push(Heap, stations[i][1])
	}
	for total < target {
		if Heap.Len() == 0 {
			return -1
		}
		total += heap.Pop(Heap).(int)
		res++
	}
	return res
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
