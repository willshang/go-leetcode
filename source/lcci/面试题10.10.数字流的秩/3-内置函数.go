package main

import "sort"

type StreamRank struct {
	arr []int
}

func Constructor() StreamRank {
	return StreamRank{arr: make([]int, 0)}
}

func (this *StreamRank) Track(x int) {
	index := sort.Search(len(this.arr), func(i int) bool {
		return x <= this.arr[i]
	})
	temp := append([]int{}, this.arr[:index]...)
	temp = append(temp, x)
	temp = append(temp, this.arr[index:]...)
	this.arr = temp
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	return sort.Search(len(this.arr), func(i int) bool {
		return x < this.arr[i]
	})
}
