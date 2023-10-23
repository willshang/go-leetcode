package main

import (
	"container/heap"
	"sort"
)

func main() {

}

type Node struct {
	Index      int
	ArriveTime int
	LeaveTime  int
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	arr := make([]Node, 0)
	waitHeap := make(WaitHeap, 0)
	heap.Init(&waitHeap)
	runHeap := make(RunHeap, 0)
	heap.Init(&runHeap)
	for i := 0; i < n; i++ {
		heap.Push(&waitHeap, i)
		arr = append(arr, Node{
			Index:      i,
			ArriveTime: times[i][0],
			LeaveTime:  times[i][1],
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].ArriveTime < arr[j].ArriveTime
	})
	cur := 0
	for i := 0; i < n; i++ {
		cur = arr[i].ArriveTime
		for runHeap.Len() > 0 && runHeap[0].EndTime <= cur { // 结束时间小于当前时间，出堆
			node := heap.Pop(&runHeap).(RunNode)
			heap.Push(&waitHeap, node.Id)
		}
		node := heap.Pop(&waitHeap).(int)
		heap.Push(&runHeap, RunNode{
			Id:      node,
			EndTime: arr[i].LeaveTime,
		})
		if arr[i].Index == targetFriend { // 目标值
			return node
		}
	}
	return -1
}

type WaitHeap []int // 空闲堆

func (h WaitHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h WaitHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h WaitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WaitHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *WaitHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

type RunNode struct {
	Id      int
	EndTime int
}

type RunHeap []RunNode // 运行堆

func (h RunHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h RunHeap) Less(i, j int) bool {
	return h[i].EndTime < h[j].EndTime
}

func (h RunHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *RunHeap) Push(x interface{}) {
	*h = append(*h, x.(RunNode))
}

func (h *RunHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
