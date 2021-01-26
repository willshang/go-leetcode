package main

import (
	"container/heap"
)

func main() {

}

func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				matrix[i][j] = matrix[i][j-1] ^ matrix[i][j]
			} else if i > 0 && j == 0 {
				matrix[i][j] = matrix[i-1][j] ^ matrix[i][j]
			} else if i > 0 && j > 0 {
				matrix[i][j] = matrix[i-1][j] ^ matrix[i][j-1] ^ matrix[i][j] ^ matrix[i-1][j-1]
			}
			heap.Push(&intHeap, matrix[i][j])
			if intHeap.Len() > k {
				heap.Pop(&intHeap)
			}
		}
	}
	return intHeap[0]
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
