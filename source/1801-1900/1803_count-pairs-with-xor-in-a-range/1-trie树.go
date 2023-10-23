package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{1, 4, 2, 7}, 2, 6))
}

// leetcode1803_统计异或值在范围内的数对有多少
func countPairs(nums []int, low int, high int) int {
	res := 0
	n := len(nums)
	root := &Trie{next: make([]*Trie, 2)}
	for i := 0; i < n; i++ {
		// 先查询，再插入
		res = res + root.Search(nums[i], high+1) - root.Search(nums[i], low)
		root.Insert(nums[i])
	}
	return res
}

type Trie struct {
	next []*Trie // 0或者1
	size int     // 次数
}

// 插入num
func (t *Trie) Insert(num int) {
	temp := t
	for i := 31; i >= 0; i-- {
		value := (num >> i) & 1
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next: make([]*Trie, 2),
			}
		}
		temp = temp.next[value]
		temp.size++
	}
}

// 查找小于target的数量
func (t *Trie) Search(num int, target int) int {
	res := 0
	temp := t
	for i := 31; i >= 0; i-- {
		if temp == nil { // 直接返回
			return res
		}
		value := (num >> i) & 1
		targetValue := (target >> i) & 1
		if targetValue > 0 { // target该位为1
			if temp.next[value] != nil {
				res = res + temp.next[value].size
			}
			temp = temp.next[1-value] // value ^ (1-value) = 1 => 往1-value走
		} else {
			temp = temp.next[value] // value ^ value = 0 // 往value走
		}
	}
	return res
}
