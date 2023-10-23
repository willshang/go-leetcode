package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(arr, 2))
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	nodeHeap := &Heap{}
	heap.Init(nodeHeap)
	for key, value := range m {
		heap.Push(nodeHeap, Node{
			str: key,
			num: value,
		})
	}
	fmt.Println(nodeHeap)
	var res []string
	for i := 0; i < k; i++ {
		value := heap.Pop(nodeHeap).(Node)
		res = append(res, value.str)
	}
	return res
}

type Node struct {
	str string
	num int
}

type Heap []Node

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	if h[i].num == h[j].num {
		return h[i].str < h[j].str
	}
	return h[i].num > h[j].num
}

func (h *Heap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
