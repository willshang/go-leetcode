package main

import "container/heap"

func main() {

}

var mod = 1000000007

var res int

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	res = target * inc // 最坏的情况：全+1
	n := len(jump)
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	visited := make(map[int]bool)
	heap.Push(&intHeap, []int{0, target}) // 时间+当前位置：从后往前走
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([]int)
		t, cur := node[0], node[1]
		if t >= res { // 跳过
			continue
		}
		res = min(res, t+cur*inc) // 用+1补
		for i := 0; i < n; i++ {
			diff, next := cur%jump[i], cur/jump[i]
			if diff == 0 { // 直接坐公交
				if visited[next] == false {
					heap.Push(&intHeap, []int{t + cost[i], next})
				}
			} else {
				if visited[next] == false { // 向左走坐公交
					heap.Push(&intHeap, []int{t + cost[i] + diff*inc, next})
				}
				if visited[next+1] == false { // 向右走坐公交
					heap.Push(&intHeap, []int{t + cost[i] + (jump[i]-diff)*dec, next + 1})
				}
			}
		}
	}
	return res % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
