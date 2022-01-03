package main

import (
	"math"
)

func main() {

}

func maximizeXor(nums []int, queries [][]int) []int {
	n := len(queries)
	root := &Trie{next: make([]*Trie, 2), minValue: math.MaxInt32}
	for i := 0; i < len(nums); i++ {
		root.Insert(nums[i])
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		target, mi := queries[i][0], queries[i][1]
		res[i] = root.getMaxValueWithLimit(target, mi)
	}
	return res
}

type Trie struct {
	next     []*Trie // 0或者1
	minValue int     // 小于当前节点的最小值
}

// 插入num
func (t *Trie) Insert(num int) {
	temp := t
	if num < temp.minValue {
		temp.minValue = num
	}
	for i := 31; i >= 0; i-- {
		value := (num >> i) & 1
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:     make([]*Trie, 2),
				minValue: num,
			}
		}
		temp = temp.next[value]
		if num < temp.minValue {
			temp.minValue = num
		}
	}
}

func (t *Trie) getMaxValueWithLimit(num int, limit int) int {
	res := 0
	temp := t
	if temp.minValue > limit {
		return -1
	}
	for i := 31; i >= 0; i-- {
		value := (num >> i) & 1
		if temp.next[1-value] != nil && temp.next[1-value].minValue <= limit { // 能取到1
			res = res | (1 << i) // 结果在该位可以为1，该为置为1
			temp = temp.next[1-value]
		} else {
			temp = temp.next[value]
		}
	}
	return res
}
