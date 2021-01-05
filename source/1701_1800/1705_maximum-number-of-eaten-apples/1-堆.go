package main

import "container/heap"

func main() {

}

func eatenApples(apples []int, days []int) int {
	res := 0
	nodeHeap := make(NodeHeap, 0)
	heap.Init(&nodeHeap)
	for i := 0; i < len(apples) || nodeHeap.Len() > 0; i++ {
		if i < len(apples) && apples[i] > 0 {
			heap.Push(&nodeHeap, Node{
				date: days[i] + i,
				num:  apples[i],
			})
		}
		for nodeHeap.Len() > 0 && nodeHeap[0].date == i {
			heap.Pop(&nodeHeap)
		}
		if nodeHeap.Len() > 0 && nodeHeap[0].num > 0 {
			res++
			nodeHeap[0].num--
			if nodeHeap[0].num == 0 {
				heap.Pop(&nodeHeap)
			}
		}
	}
	return res
}

type Node struct {
	date int
	num  int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h NodeHeap) Less(i, j int) bool {
	return h[i].date < h[j].date
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
