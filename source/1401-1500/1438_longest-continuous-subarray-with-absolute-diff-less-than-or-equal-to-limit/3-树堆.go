package main

import (
	"math/rand"
)

func longestSubarray(nums []int, limit int) int {
	res := 0
	t := &treap{}
	var i int
	for j := 0; j < len(nums); j++ {
		t.insert(nums[j])
		for t.max().key-t.min().key > limit {
			t.delete(nums[i])
			i++
		}
		res = max(res, j-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	ch       [2]*Node
	priority int
	key      int
	cnt      int
}

func (n *Node) cmp(b int) int {
	switch {
	case b < n.key:
		return 0
	case b > n.key:
		return 1
	default:
		return -1
	}
}

func (n *Node) rotate(b int) *Node {
	target := n.ch[b^1]
	n.ch[b^1] = target.ch[b]
	target.ch[b] = n
	return target
}

type treap struct {
	root *Node
}

func (t *treap) ins(n *Node, key int) *Node {
	if n == nil {
		return &Node{
			ch:       [2]*Node{}, // 0左边，1右边
			priority: rand.Int(),
			key:      key,
			cnt:      1,
		}
	}
	if b := n.cmp(key); b >= 0 {
		n.ch[b] = t.ins(n.ch[b], key)
		if n.ch[b].priority > n.priority {
			n = n.rotate(b ^ 1)
		}
	} else {
		n.cnt++
	}
	return n
}

func (t *treap) del(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if b := n.cmp(key); b >= 0 {
		n.ch[b] = t.del(n.ch[b], key)
	} else {
		if n.cnt > 1 {
			n.cnt--
		} else {
			if n.ch[1] == nil {
				return n.ch[0]
			}
			if n.ch[0] == nil {
				return n.ch[1]
			}
			b = 0
			if n.ch[0].priority > n.ch[1].priority {
				b = 1
			}
			n = n.rotate(b)
			n.ch[b] = t.del(n.ch[b], key)
		}
	}
	return n
}

func (t *treap) insert(key int) {
	t.root = t.ins(t.root, key)
}

func (t *treap) delete(key int) {
	t.root = t.del(t.root, key)
}

func (t *treap) min() (min *Node) {
	for temp := t.root; temp != nil; temp = temp.ch[0] {
		min = temp
	}
	return min
}

func (t *treap) max() (max *Node) {
	for temp := t.root; temp != nil; temp = temp.ch[1] {
		max = temp
	}
	return max
}
