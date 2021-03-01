package main

import (
	"math/rand"
	"time"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode382_链表随机节点
type Solution struct {
	head *ListNode
	r    *rand.Rand
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
		r:    rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (this *Solution) GetRandom() int {
	res := this.head.Val
	cur := this.head
	count := 1
	for cur != nil {
		if this.r.Intn(count)+1 == count {
			res = cur.Val
		}
		count++
		cur = cur.Next
	}
	return res
}
