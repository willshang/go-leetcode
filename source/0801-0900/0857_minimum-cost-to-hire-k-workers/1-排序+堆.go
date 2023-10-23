package main

import (
	"container/heap"
	"math"
	"sort"
)

func main() {

}

// leetcode857_雇佣K名工人的最低成本
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		arr[i] = [2]int{quality[i], wage[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		a := float64(arr[i][1]) / float64(arr[i][0])
		b := float64(arr[j][1]) / float64(arr[j][0])
		return a < b
	})
	res := math.MaxFloat64
	var sum float64
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	// 枚举最低时薪或者性价比，然后在能接受最低时薪的人中选择工作质量总和最小的k个人
	for i := 0; i < n; i++ {
		cur := float64(arr[i][1]) / float64(arr[i][0]) // 性价比：最低工资/质量， 比率从小到大遍历
		heap.Push(&intHeap, arr[i][0])                 // 质量高优先淘汰
		sum = sum + float64(arr[i][0])                 // 质量和
		if intHeap.Len() > k {
			node := heap.Pop(&intHeap).(int)
			sum = sum - float64(node)
		}
		if intHeap.Len() == k && cur*sum < res {
			res = cur * sum
		}
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
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
