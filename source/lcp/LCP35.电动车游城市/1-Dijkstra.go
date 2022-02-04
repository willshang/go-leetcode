package main

import (
	"container/heap"
	"math"
)

func main() {

}

// leetcode_lcp35_电动车游城市
func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	arr := make([][][]int, n) // 邻接表
	for i := 0; i < len(paths); i++ {
		a, b, c := paths[i][0], paths[i][1], paths[i][2] // a<=>b
		arr[a] = append(arr[a], []int{b, c})
		arr[b] = append(arr[b], []int{a, c})
	}
	dis := make([][]int, n) // start到i点在j电量下的花费
	for i := 0; i < n; i++ {
		dis[i] = make([]int, cnt+1)
		for j := 0; j < cnt+1; j++ {
			dis[i][j] = math.MaxInt32
		}
	}
	dis[start][0] = 0 // 开始花费为0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, []int{0, start, 0}) // 时间+位置+电量：按时间堆排序
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([]int)
		t, cur, value := node[0], node[1], node[2]
		if t > dis[cur][value] { // 大于跳过
			continue
		}
		if cur == end { // 终点，直接返回
			return t
		}
		// 核心点：可以在一个城市充满电，可以不需要在其它城市充电；
		// 这样可以尝试在一个城市充满一定电量后往下走；不一定非要完全充满电或者只充到达下一个城市所需的电量
		// 因为每个城市的充电单价不一样
		if value < cnt { // 去第cur城市充电1单位；入堆后可以继续充电1单位
			nextTime := t + charge[cur]
			if nextTime < dis[cur][value+1] {
				dis[cur][value+1] = nextTime
				heap.Push(&intHeap, []int{nextTime, cur, value + 1})
			}
		}
		// 电量满足到达下一个城市后可以开往下一个城市
		for i := 0; i < len(arr[cur]); i++ {
			next, nextDis := arr[cur][i][0], arr[cur][i][1]
			if value >= nextDis && t+nextDis < dis[next][value-nextDis] {
				dis[next][value-nextDis] = t + nextDis
				heap.Push(&intHeap, []int{t + nextDis, next, value - nextDis})
			}
		}
	}
	return -1
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
