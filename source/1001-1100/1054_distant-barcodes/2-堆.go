package main

import "container/heap"

func main() {

}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	if n <= 1 {
		return barcodes
	}
	res := make([]int, 0)
	m := make(map[int]int)
	for _, v := range barcodes {
		m[v]++
	}
	nodeHeap := &Heap{}
	heap.Init(nodeHeap)
	for k, v := range m {
		heap.Push(nodeHeap, Node{
			value: k,
			num:   v,
		})
	}
	for nodeHeap.Len() >= 2 {
		node1 := heap.Pop(nodeHeap).(Node)
		node2 := heap.Pop(nodeHeap).(Node)
		res = append(res, node1.value, node2.value)
		node1.num--
		node2.num--
		if node1.num > 0 {
			heap.Push(nodeHeap, node1)
		}
		if node2.num > 0 {
			heap.Push(nodeHeap, node2)
		}
	}
	if nodeHeap.Len() > 0 {
		t := heap.Pop(nodeHeap).(Node)
		res = append(res, t.value)
	}
	return res
}

type Node struct {
	value int
	num   int
}

type Heap []Node

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].num > h[j].num
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *Heap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
