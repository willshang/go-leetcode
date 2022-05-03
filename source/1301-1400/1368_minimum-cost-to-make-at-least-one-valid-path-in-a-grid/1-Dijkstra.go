package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minCost([][]int{
		{1, 1, 3},
		{3, 2, 2},
		{1, 1, 4},
	}))
}

// 右左下上
var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func minCost(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	total := n * m
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			arr[i][j] = total
		}
	}
	arr[0][0] = 0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, []int{0, 0, 0})
	visited := make(map[[2]int]bool)
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([]int)
		v, a, b := node[0], node[1], node[2]
		if visited[[2]int{a, b}] == true {
			continue
		}
		visited[[2]int{a, b}] = true
		for i := 0; i < 4; i++ {
			x, y := a+dx[i], b+dy[i]
			if 0 <= x && x < n && 0 <= y && y < m {
				dis := v
				if i+1 != grid[a][b] {
					dis++ // 变换方向+1
				}
				if dis < arr[x][y] {
					arr[x][y] = dis
					heap.Push(&intHeap, []int{dis, x, y})
				}
			}
		}
	}
	return arr[n-1][m-1]
}

type IntHeap [][]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
