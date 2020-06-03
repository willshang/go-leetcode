package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 1))
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

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	intHeap := make(IntHeap, 0, k)
	heap.Init(&intHeap)

	for i := 0; i < len(arr); i++ {
		if len(intHeap) < k {
			heap.Push(&intHeap, arr[i])
		} else if len(intHeap) == k {
			if arr[i] < intHeap[0] {
				heap.Pop(&intHeap)
				heap.Push(&intHeap, arr[i])
			}
		}
	}
	return intHeap
}
