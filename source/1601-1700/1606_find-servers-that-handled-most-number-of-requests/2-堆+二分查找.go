package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(busiestServers(3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}))
}

// leetcode1606_找到处理最多请求的服务器
func busiestServers(k int, arrival []int, load []int) []int {
	n := len(arrival)
	res := make([]int, 0)
	arr := make([]int, k)
	free := make([]int, k)
	for i := 0; i < k; i++ {
		free[i] = i
	}
	busyHeap := make(IntHeap, 0)
	heap.Init(&busyHeap)
	for i := 0; i < n; i++ {
		start := arrival[i]
		// busy堆执行完出队列
		for busyHeap.Len() > 0 && busyHeap[0][0] <= start {
			// free插入该下标：插入有序数组后有序
			id := busyHeap[0][1]
			index := sort.SearchInts(free, id)
			free = append(free[:index], append([]int{id}, free[index:]...)...)
			heap.Pop(&busyHeap)
		}
		if len(free) == 0 {
			continue
		}
		id := sort.SearchInts(free, i%k)
		if id == len(free) {
			id = 0
		}
		arr[free[id]]++
		heap.Push(&busyHeap, []int{start + load[i], free[id]})
		free = append(free[:id], free[id+1:]...) // 删除该下标
	}
	var maxValue int
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
			res = []int{i}
		} else if arr[i] == maxValue {
			res = append(res, i)
		}
	}
	return res
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
