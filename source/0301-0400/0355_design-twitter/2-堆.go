package main

import (
	"container/heap"
	"fmt"
)

func main() {
	t := Constructor()
	t.PostTweet(1, 5)
	fmt.Println(t.GetNewsFeed(1))
	t.Follow(1, 2)
	t.PostTweet(2, 6)
	fmt.Println(t.GetNewsFeed(1))
	t.Unfollow(1, 2)
	fmt.Println(t.GetNewsFeed(1))
}

type Twitter struct {
	data  map[int][][2]int
	m     map[int]map[int]int
	count int
}

func Constructor() Twitter {
	return Twitter{
		data:  make(map[int][][2]int),
		m:     make(map[int]map[int]int),
		count: 0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.data[userId] = append(this.data[userId], [2]int{this.count, tweetId})
	this.count++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0)
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := len(this.data[userId]) - 1; i >= 0; i-- {
		a, b := this.data[userId][i][0], this.data[userId][i][1]
		if intHeap.Len() == 10 && intHeap[0][0] > a {
			break
		}
		heap.Push(&intHeap, [2]int{a, b})
		if intHeap.Len() > 10 {
			heap.Pop(&intHeap)
		}
	}
	for k, v := range this.m[userId] {
		if k == userId || v == 0 {
			continue
		}
		for i := len(this.data[k]) - 1; i >= 0; i-- {
			a, b := this.data[k][i][0], this.data[k][i][1]
			if intHeap.Len() == 10 && intHeap[0][0] > a {
				break
			}
			heap.Push(&intHeap, [2]int{a, b})
			if intHeap.Len() > 10 {
				heap.Pop(&intHeap)
			}
		}
	}
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([2]int)
		res = append([]int{node[1]}, res...)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.m[followerId] == nil {
		this.m[followerId] = make(map[int]int)
	}
	this.m[followerId][followeeId] = 1
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.m[followerId] != nil {
		this.m[followerId][followeeId] = 0
	}
}

type IntHeap [][2]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
