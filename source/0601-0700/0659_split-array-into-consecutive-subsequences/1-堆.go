package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{4, 5, 6, 7, 7, 8, 8, 9, 10, 11}))
}

func isPossible(nums []int) bool {
	m := make(map[int]*IntHeap)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if m[v] == nil {
			intHeap := make(IntHeap, 0)
			heap.Init(&intHeap)
			m[v] = &intHeap
		}
		length := 0
		if h := m[v-1]; h != nil {
			length = heap.Pop(h).(int) // 找到最短的以v-1结尾的长度
			if m[v-1].Len() == 0 {
				delete(m, v-1)
			}
		}
		temp := m[v]
		heap.Push(temp, length+1)
	}
	for _, v := range m {
		if (*v)[0] < 3 {
			return false
		}
	}
	return true
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
