package main

import "container/heap"

func main() {

}

// leetcode1792_最大平均通过率
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	nodeHeap := make(NodeHeap, 0)
	heap.Init(&nodeHeap)
	for i := 0; i < len(classes); i++ {
		x, y := float64(classes[i][0]), float64(classes[i][1])
		a := x / y
		b := (x + 1) / (y + 1)
		heap.Push(&nodeHeap, Node{
			id:    i,
			ratio: b - a,
		})
	}
	for i := 0; i < extraStudents; i++ {
		node := heap.Pop(&nodeHeap).(Node)
		id := node.id
		classes[id][0]++
		classes[id][1]++
		x, y := float64(classes[id][0]), float64(classes[id][1])
		a := x / y
		b := (x + 1) / (y + 1)
		heap.Push(&nodeHeap, Node{
			id:    id,
			ratio: b - a,
		})
	}
	sum := float64(0)
	for i := 0; i < len(classes); i++ {
		x, y := float64(classes[i][0]), float64(classes[i][1])
		sum = sum + x/y
	}
	return sum / float64(len(classes))
}

type Node struct {
	id    int
	ratio float64
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h NodeHeap) Less(i, j int) bool {
	return h[i].ratio > h[j].ratio
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
