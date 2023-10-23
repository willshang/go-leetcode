package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	maxProfit := &ProfitNode{}
	minCapital := &CapitalNode{}
	heap.Init(maxProfit)
	heap.Init(minCapital)
	for i := 0; i < len(Profits); i++ {
		heap.Push(minCapital, Node{
			Profits: Profits[i],
			Capital: Capital[i],
		})
	}
	for i := 0; i < k; i++ {
		for minCapital.Len() > 0 {
			node := heap.Pop(minCapital).(Node)
			if node.Capital <= W {
				heap.Push(maxProfit, node)
			} else {
				heap.Push(minCapital, node)
				break
			}
		}
		if maxProfit.Len() == 0 {
			return W
		}
		node := heap.Pop(maxProfit).(Node)
		W = W + node.Profits
	}
	return W
}

type Node struct {
	Profits int
	Capital int
}

type ProfitNode []Node

func (h ProfitNode) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h ProfitNode) Less(i, j int) bool {
	return h[i].Profits > h[j].Profits
}

func (h ProfitNode) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ProfitNode) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *ProfitNode) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

type CapitalNode []Node

func (h CapitalNode) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h CapitalNode) Less(i, j int) bool {
	return h[i].Capital < h[j].Capital
}

func (h CapitalNode) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CapitalNode) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *CapitalNode) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
