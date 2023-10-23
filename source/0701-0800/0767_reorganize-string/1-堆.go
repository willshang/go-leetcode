package main

import "container/heap"

func main() {

}

func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}
	res := make([]byte, 0)
	m := make(map[byte]int)
	for _, v := range S {
		m[byte(v)]++
	}
	nodeHeap := &Heap{}
	heap.Init(nodeHeap)
	for k, v := range m {
		if v > (n+1)/2 {
			return ""
		}
		heap.Push(nodeHeap, Node{
			char: k,
			num:  v,
		})
	}
	for nodeHeap.Len() >= 2 {
		node1 := heap.Pop(nodeHeap).(Node)
		node2 := heap.Pop(nodeHeap).(Node)
		res = append(res, node1.char, node2.char)
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
		res = append(res, t.char)
	}
	return string(res)
}

type Node struct {
	char byte
	num  int
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
