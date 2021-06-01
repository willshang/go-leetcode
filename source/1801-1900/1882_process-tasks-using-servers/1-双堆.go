package main

import "container/heap"

func main() {

}

// leetcode1882_使用服务器处理任务
func assignTasks(servers []int, tasks []int) []int {
	res := make([]int, 0)
	n := len(servers)
	waitHeap := make(WaitHeap, 0)
	heap.Init(&waitHeap)
	runHeap := make(RunHeap, 0)
	heap.Init(&runHeap)
	for i := 0; i < n; i++ {
		heap.Push(&waitHeap, Node{
			Id:   i,
			Rank: servers[i],
		})
	}
	curTime := 0
	for i := 0; i < len(tasks); i++ {
		curTime = max(curTime, i)
		if waitHeap.Len() == 0 { // 无服务器可用，时间移动
			endTime := runHeap[0].EndTime // 最小的结束时间的出来
			for runHeap.Len() > 0 && runHeap[0].EndTime == endTime {
				node := heap.Pop(&runHeap).(Node)
				heap.Push(&waitHeap, node)
			}
			curTime = max(curTime, endTime)
		} else {
			for runHeap.Len() > 0 && runHeap[0].EndTime <= curTime {
				node := heap.Pop(&runHeap).(Node)
				heap.Push(&waitHeap, node)
			}
		}
		node := heap.Pop(&waitHeap).(Node)
		res = append(res, node.Id)
		heap.Push(&runHeap, Node{
			Id:      node.Id,
			Rank:    node.Rank,
			EndTime: curTime + tasks[i],
		})
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Id      int
	Rank    int // 权重
	EndTime int
}

// 运行堆
type RunHeap []Node

func (h RunHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h RunHeap) Less(i, j int) bool {
	return h[i].EndTime < h[j].EndTime // 按照时间排序
}

func (h RunHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *RunHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *RunHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

// 等待堆
type WaitHeap []Node

func (h WaitHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h WaitHeap) Less(i, j int) bool {
	if h[i].Rank == h[j].Rank {
		return h[i].Id < h[j].Id
	}
	return h[i].Rank < h[j].Rank
}

func (h WaitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WaitHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *WaitHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
