package main

import "container/heap"

func main() {

}

type mixHeap struct {
	arr   [][]int
	isBig bool
}

func (m *mixHeap) Len() int {
	return len(m.arr)
}

func (m *mixHeap) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func (m *mixHeap) Less(i, j int) bool {
	if m.isBig {
		return m.arr[i][0] > m.arr[j][0] // 大根堆
	}
	return m.arr[i][0] < m.arr[j][0] // 小根堆
}

func (m *mixHeap) Push(x interface{}) {
	m.arr = append(m.arr, x.([]int))
}

func (m *mixHeap) Pop() interface{} {
	value := (m.arr)[len(m.arr)-1]
	m.arr = (m.arr)[:len(m.arr)-1]
	return value
}

func (m *mixHeap) push(x []int) {
	heap.Push(m, x)
}

func (m *mixHeap) pop() []int {
	return heap.Pop(m).([]int)
}

func (m *mixHeap) Top() []int {
	if m.Len() > 0 {
		return m.arr[0]
	}
	return nil
}
