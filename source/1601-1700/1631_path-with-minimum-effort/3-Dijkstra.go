package main

import (
	"container/heap"
	"math"
)

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	arr := make([][]int, n) //  高度差绝对值的最大值
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			arr[i][j] = math.MaxInt32 // 初始化为最大值
		}
	}
	arr[0][0] = 0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [3]int{0, 0, 0})
	for {
		node := heap.Pop(&intHeap).([3]int)
		a, b, c := node[0], node[1], node[2]
		if a == n-1 && b == m-1 { // 返回结果
			return c
		}
		if arr[a][b] < c { // 大于 高度差绝对值的最大值，跳过
			continue
		}
		for i := 0; i < 4; i++ {
			x, y := a+dx[i], b+dy[i]
			if 0 <= x && x < n && 0 <= y && y < m {
				diff := max(c, abs(heights[x][y]-heights[a][b])) // 更新：去 高度差绝对值的最大值 的较大值
				if diff < arr[x][y] {                            // 更新：加入堆，点会重复，通过更新arr[x][y]来过滤
					arr[x][y] = diff
					heap.Push(&intHeap, [3]int{x, y, diff})
				}
			}
		}
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
