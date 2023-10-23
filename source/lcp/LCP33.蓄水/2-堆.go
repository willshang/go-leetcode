package main

import (
	"container/heap"
	"math"
)

func main() {

}

func storeWater(bucket []int, vat []int) int {
	n := len(vat)
	nodeHeap := make(IntHeap, 0)
	heap.Init(&nodeHeap)
	count := 0 // 需要升级的次数
	for i := 0; i < n; i++ {
		if bucket[i] == 0 && vat[i] > 0 {
			bucket[i] = 1
			count++
		}
		if vat[i] > 0 {
			heap.Push(&nodeHeap, Node{
				bucket: bucket[i],
				vat:    vat[i],
				count:  (vat[i]-1)/bucket[i] + 1,
			})
		}
	}
	res := math.MaxInt32 // 总次数
	for nodeHeap.Len() > 0 {
		node := heap.Pop(&nodeHeap).(Node)
		if count >= res {
			break
		}
		res = min(res, node.count+count) // 堆里面最大的蓄水次数+升级的次数
		heap.Push(&nodeHeap, Node{
			bucket: node.bucket + 1,
			vat:    node.vat,
			count:  (node.vat-1)/(node.bucket+1) + 1,
		})
		count++
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type Node struct {
	bucket int
	vat    int
	count  int
}

type IntHeap []Node

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
