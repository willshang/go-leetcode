# 堆

## go-使用接口实现
- 堆的接口在`"str/container/heap/heap.go"`
```go
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}
```

```go
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
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

func heapSort(arr []int) []int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(arr); i++ {
		heap.Push(&intHeap, arr[i])
	}
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		value := heap.Pop(&intHeap).(int)
		res = append(res, value)
	}
	return res
}

func main() {
	arr := []int{1, 10, 9, 5, 6, 7, 4, 8, 2, 3}
	fmt.Println(heapSort(arr))
}
```