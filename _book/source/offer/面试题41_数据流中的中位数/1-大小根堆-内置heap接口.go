package main

import "container/heap"

func main() {

}

// 剑指offer_面试题41_数据流中的中位数
type MinHeap []int

func (i MinHeap) Len() int {
	return len(i)
}

func (i MinHeap) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i MinHeap) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}
func (i *MinHeap) Push(v interface{}) {
	*i = append(*i, v.(int))
}

func (i *MinHeap) Pop() interface{} {
	value := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	return value
}

type MaxHeap []int

func (i MaxHeap) Len() int {
	return len(i)
}

func (i MaxHeap) Less(x, y int) bool {
	return i[x] > i[y]
}

func (i MaxHeap) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i *MaxHeap) Push(v interface{}) {
	*i = append(*i, v.(int))
}

func (i *MaxHeap) Pop() interface{} {
	value := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	return value
}

type MedianFinder struct {
	minArr *MinHeap
	maxArr *MaxHeap
}

func Constructor() MedianFinder {
	res := new(MedianFinder)
	res.minArr = new(MinHeap)
	res.maxArr = new(MaxHeap)
	heap.Init(res.minArr)
	heap.Init(res.maxArr)
	return *res
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxArr.Len() == this.minArr.Len() {
		heap.Push(this.minArr, num)
		heap.Push(this.maxArr, heap.Pop(this.minArr))
	} else {
		heap.Push(this.maxArr, num)
		heap.Push(this.minArr, heap.Pop(this.maxArr))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minArr.Len() == this.maxArr.Len() {
		return (float64((*this.maxArr)[0]) + float64((*this.minArr)[0])) / 2
	} else {
		return float64((*this.maxArr)[0])
	}
}
