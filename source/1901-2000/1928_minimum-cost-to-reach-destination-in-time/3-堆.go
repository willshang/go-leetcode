package main

import (
	"container/heap"
)

func main() {

}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	arr := make([][][2]int, n)
	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1], edges[i][2] // a=>b c
		arr[a] = append(arr[a], [2]int{b, c})
		arr[b] = append(arr[b], [2]int{a, c})
	}
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [3]int{passingFees[0], maxTime, 0}) // 费用，剩余时间，下标
	m := make(map[int]int)
	m[0] = maxTime
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([3]int)
		value, leftTime, index := node[0], node[1], node[2]
		if index == n-1 {
			return value
		}
		for i := 0; i < len(arr[index]); i++ {
			a, b := arr[index][i][0], arr[index][i][1]
			if b > leftTime { // 大于剩余时间
				continue
			}
			if v, ok := m[a]; ok == false || leftTime-b > v {
				m[a] = leftTime - b
				heap.Push(&intHeap, [3]int{value + passingFees[a], leftTime - b, a})
			}
		}
	}
	return -1
}

type IntHeap [][3]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
