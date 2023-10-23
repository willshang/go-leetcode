package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	arr := make([][]Node, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], Node{index: b, Value: succProb[i]})
		arr[b] = append(arr[b], Node{index: a, Value: succProb[i]})
	}
	visited := make([]bool, n)
	maxValue := make([]float64, n)
	nodeHeap := make(NodeHeap, 0)
	heap.Init(&nodeHeap)
	heap.Push(&nodeHeap, Node{
		Value: 1,
		index: start,
	})
	for nodeHeap.Len() > 0 {
		node := heap.Pop(&nodeHeap).(Node)
		visited[node.index] = true
		if node.index == end {
			return node.Value
		}
		for i := 0; i < len(arr[node.index]); i++ {
			next := arr[node.index][i]
			if visited[next.index] == true || node.Value*next.Value < maxValue[next.index] {
				continue
			}
			maxValue[next.index] = node.Value * next.Value
			heap.Push(&nodeHeap, Node{
				Value: maxValue[next.index],
				index: next.index,
			})
		}
	}
	return 0
}

type Node struct {
	Value float64
	index int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h NodeHeap) Less(i, j int) bool {
	return h[i].Value > h[j].Value
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
