package main

import (
	"container/heap"
	"math"
)

func main() {

}

func smallestRange(nums [][]int) []int {
	nodeHeap := make(NodeHeap, 0) // 堆的大小为n
	heap.Init(&nodeHeap)
	maxValue, n := math.MinInt32, len(nums)
	// 问题可以转化为，从n个列表中各取1个数，使得这n个数中的最大值与最小值的差最小
	for i := 0; i < n; i++ {
		maxValue = max(maxValue, nums[i][0]) // 获取n个数的最大值
		heap.Push(&nodeHeap, Node{Id: i, Value: nums[i][0]})
		nums[i] = nums[i][1:] // 数组缩小，也可以使用下标标记
	}
	res := []int{math.MinInt32 / 10, math.MaxInt32 / 10}
	for { // 从小到大，每从堆取出一个最小值，再从所在组取出下一个较大的数放回去
		node := heap.Pop(&nodeHeap).(Node)       // 小根堆，取最小值
		if maxValue-node.Value < res[1]-res[0] { // 更新范围：最大值-最小值
			res = []int{node.Value, maxValue}
		}
		if len(nums[node.Id]) == 0 { // 退出条件：某一个数组首先访问完
			break
		}
		heap.Push(&nodeHeap, Node{Id: node.Id, Value: nums[node.Id][0]})
		maxValue = max(maxValue, nums[node.Id][0]) // 更新最大值
		nums[node.Id] = nums[node.Id][1:]
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
	Id    int
	Value int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h NodeHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
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
