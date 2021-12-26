package main

import "container/heap"

func main() {

}

// leetcode2034_股票价格波动
type StockPrice struct {
	maxH, minH        *mixHeap
	curTime, curPrice int
	m                 map[int]int
}

func Constructor() StockPrice {
	return StockPrice{
		maxH: &mixHeap{isBig: true},
		minH: &mixHeap{isBig: false},
		m:    make(map[int]int), // 保存最新的时间戳=>价格数据
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if timestamp >= this.curTime { // 更新
		this.curTime, this.curPrice = timestamp, price
	}
	this.maxH.push([]int{timestamp, price})
	this.minH.push([]int{timestamp, price})
	this.m[timestamp] = price // 时间戳=>价格
}

func (this *StockPrice) Current() int {
	return this.curPrice
}

func (this *StockPrice) Maximum() int {
	for this.maxH.Len() > 0 {
		top := this.maxH.Top()
		if top[1] == this.m[top[0]] { // 跟map里面的价格对的上，直接返回
			return top[1]
		}
		this.maxH.pop() // 剔除旧数据
	}
	return 0
}

func (this *StockPrice) Minimum() int {
	for this.minH.Len() > 0 {
		top := this.minH.Top()
		if top[1] == this.m[top[0]] { // 跟map里面的价格对的上，直接返回
			return top[1]
		}
		this.minH.pop() // 剔除旧数据
	}
	return 0
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
		return m.arr[i][1] > m.arr[j][1]
	}
	return m.arr[i][1] < m.arr[j][1]
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
