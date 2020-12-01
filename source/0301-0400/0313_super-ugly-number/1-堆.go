package main

import "container/heap"

func main() {

}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 0 || n == 1 {
		return n
	}
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	heap.Push(intHeap, 1)
	n--
	for n > 0 {
		x := heap.Pop(intHeap).(int)
		for intHeap.Len() > 0 && x == (*intHeap)[0] {
			heap.Pop(intHeap)
		}
		for i := 0; i < len(primes); i++ {
			heap.Push(intHeap, x*primes[i])
		}
		n--
	}
	return heap.Pop(intHeap).(int)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

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
