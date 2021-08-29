package main

import "container/heap"

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func swimInWater(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [3]int{0, 0, grid[0][0]})
	res := 0
	for {
		node := heap.Pop(&intHeap).([3]int)
		a, b, c := node[0], node[1], node[2]
		res = max(res, c) // 更新时间
		// 经过时间t以后，可以瞬间从坐标[0,0]到坐标[N-1, N-1]。
		if a == n-1 && b == n-1 {
			return res
		}
		for i := 0; i < 4; i++ {
			x, y := a+dx[i], b+dy[i]
			if 0 <= x && x < n && 0 <= y && y < n && visited[x][y] == false {
				visited[x][y] = true
				heap.Push(&intHeap, [3]int{x, y, grid[x][y]})
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type IntHeap [][3]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
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
