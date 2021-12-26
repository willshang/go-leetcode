package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
}

type StockPrice struct {
	*redblacktree.Tree
	m                 map[int]int
	curTime, curPrice int
}

func Constructor() StockPrice {
	return StockPrice{
		Tree: redblacktree.NewWithIntComparator(),
		m:    make(map[int]int),
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if v, ok := this.m[timestamp]; ok {
		this.remove(v) // 删除
	}
	this.put(price) // 添加记录
	this.m[timestamp] = price
	if timestamp >= this.curTime { // 更新
		this.curTime, this.curPrice = timestamp, price
	}
}

func (this *StockPrice) put(price int) {
	count := 0
	if v, ok := this.Get(price); ok {
		count = v.(int)
	}
	this.Put(price, count+1)
}

func (this *StockPrice) remove(price int) {
	if v, ok := this.Get(price); ok && v.(int) > 1 {
		this.Put(price, v.(int)-1)
	} else {
		this.Remove(price)
	}
}

func (this *StockPrice) Current() int {
	return this.curPrice
}

func (this *StockPrice) Maximum() int {
	return this.Right().Key.(int)
}

func (this *StockPrice) Minimum() int {
	return this.Left().Key.(int)
}
