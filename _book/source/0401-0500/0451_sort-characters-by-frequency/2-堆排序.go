package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(frequencySort("tree"))
}

func frequencySort(s string) string {
	m := make(map[byte]string)
	for i := 0; i < len(s); i++ {
		m[s[i]] = m[s[i]] + string(s[i])
	}
	var h HeapString
	heap.Init(&h)
	for _, v := range m {
		heap.Push(&h, v)
	}
	res := ""
	for h.Len() > 0 {
		str := heap.Pop(&h).(string)
		res = res + str
	}
	return res
}

type HeapString []string

func (h HeapString) Len() int {
	return len(h)
}

func (h HeapString) Less(i int, j int) bool {
	return len(h[i]) >= len(h[j])
}

func (h HeapString) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapString) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *HeapString) Pop() interface{} {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[:n-1]
	return val
}
