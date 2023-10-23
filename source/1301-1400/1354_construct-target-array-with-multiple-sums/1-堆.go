package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(isPossible([]int{5, 8}))
	fmt.Println(isPossible([]int{1, 1, 3}))
}

// leetcode1354_多次求和构造目标数组
func isPossible(target []int) bool {
	sum := 0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(target); i++ {
		sum = sum + target[i]
		heap.Push(&intHeap, target[i])
	}
	for {
		curMax := heap.Pop(&intHeap).(int) // 当前轮最大值（从大到小替换）
		if curMax == 1 {                   // 全是1的情况
			break
		}
		otherSum := sum - curMax
		// [5,8] 8 13
		// [3,5] 5 8
		// [2,3] 3 5
		// [1,2] 2 3
		if curMax <= otherSum || otherSum == 0 { // 例如[1,1,2]=>curMax=2, otherSum=2的情况
			return false
		}
		temp := curMax % otherSum
		heap.Push(&intHeap, temp)
		sum = sum - curMax + temp
	}
	return true
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
