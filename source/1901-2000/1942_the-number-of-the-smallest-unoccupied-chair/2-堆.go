package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode1942_最小未被占据椅子的编号
func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	waitHeap := make(WaitHeap, 0)
	heap.Init(&waitHeap)
	arriveArr := make([][2]int, 0)
	leaveArr := make([][2]int, 0)
	for i := 0; i < n; i++ {
		heap.Push(&waitHeap, i)
		arriveArr = append(arriveArr, [2]int{i, times[i][0]})
		leaveArr = append(leaveArr, [2]int{i, times[i][1]})
	}
	sort.Slice(arriveArr, func(i, j int) bool {
		return arriveArr[i][1] < arriveArr[j][1]
	})
	sort.Slice(leaveArr, func(i, j int) bool {
		return leaveArr[i][1] < leaveArr[j][1]
	})

	j := 0
	m := make(map[int]int) // 人=>座位
	for i := 0; i < n; i++ {
		for j < n && leaveArr[j][1] <= arriveArr[i][1] { // 小于当前时间：出堆
			heap.Push(&waitHeap, m[leaveArr[j][0]])
			j++
		}
		target := heap.Pop(&waitHeap).(int)
		m[arriveArr[i][0]] = target
		if arriveArr[i][0] == targetFriend { // 目标值
			return target
		}
	}
	return -1
}

type WaitHeap []int

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
