package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{}, 0))
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	intHeap := make(IntHeap, 0, k)
	heap.Init(&intHeap)
	for i := 0; i < k; i++ {
		heap.Push(&intHeap, nums[i])
	}
	for i := k; i < len(nums); i++ {
		temp := heap.Pop(&intHeap).(int)
		res = append(res, temp)
		if temp != nums[i-k] {
			intHeap.Remove(nums[i-k])
			heap.Push(&intHeap, temp)
			heap.Push(&intHeap, nums[i])
		} else {
			heap.Push(&intHeap, nums[i])
		}
	}
	res = append(res, heap.Pop(&intHeap).(int))
	return res
}

type IntHeap []int

func (i IntHeap) Len() int {
	return len(i)
}

func (i IntHeap) Less(x, y int) bool {
	return i[x] > i[y]
}

func (i IntHeap) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}
func (i *IntHeap) Push(v interface{}) {
	*i = append(*i, v.(int))
}

func (i *IntHeap) Pop() interface{} {
	value := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	return value
}

func (i *IntHeap) Remove(x interface{}) {
	for j := 0; j < len(*i); j++ {
		if (*i)[j] == x {
			*i = append((*i)[:j], (*i)[j+1:]...)
			break
		}
	}
	heap.Init(i)
}
