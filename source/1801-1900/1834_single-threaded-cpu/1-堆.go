package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode1834_单线程CPU
func getOrder(tasks [][]int) []int {
	n := len(tasks)
	res := make([]int, 0)
	arr := make([]Node, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, Node{
			Id:             i,
			StartTime:      tasks[i][0],
			ProcessingTime: tasks[i][1],
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].StartTime < arr[j].StartTime
	})
	nodeHeap := make(NodeHeap, 0)
	heap.Init(&nodeHeap)
	curTime := 0
	cur := 0
	for i := 0; i < n; i++ { // 每次处理1个任务
		if nodeHeap.Len() == 0 { // 空任务
			curTime = max(curTime, tasks[arr[cur].Id][0]) // 时间移动
		}
		for ; cur < n && arr[cur].StartTime <= curTime; cur++ { // 加入优先队列：将小于等于时间戳的任务加入堆
			heap.Push(&nodeHeap, arr[cur])
		}
		node := heap.Pop(&nodeHeap).(Node)      // 选择任务：选择处理时间小的任务
		curTime = curTime + node.ProcessingTime // 时间处理
		res = append(res, node.Id)
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
	Id             int
	StartTime      int
	ProcessingTime int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h NodeHeap) Less(i, j int) bool {
	if h[i].ProcessingTime == h[j].ProcessingTime {
		return h[i].Id < h[j].Id
	}
	return h[i].ProcessingTime < h[j].ProcessingTime
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
