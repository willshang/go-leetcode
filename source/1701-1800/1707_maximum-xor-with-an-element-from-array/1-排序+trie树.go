package main

import "sort"

func main() {

}

// leetcode1707_与数组中元素的最大异或值
func maximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	n := len(queries)
	for i := 0; i < n; i++ {
		queries[i] = append(queries[i], i) // 每个查询添加1个下标方便排序后找到原来的位置
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1] < queries[j][1] // 按照mi从小到大排序
	})
	root := &Trie{next: make([]*Trie, 2)}
	res := make([]int, n)
	var j int
	for i := 0; i < n; i++ {
		target, mi, index := queries[i][0], queries[i][1], queries[i][2]
		for j < len(nums) && nums[j] <= mi { // 插入小于等于mi的数
			root.Insert(nums[j])
			j++
		}
		if j == 0 {
			res[index] = -1
		} else {
			res[index] = root.getMaxValue(target)
		}
	}
	return res
}

type Trie struct {
	next []*Trie // 0或者1
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
	}
}

// 查找异或对应的最大值
func (t *Trie) getMaxValue(num int) int {
	res := 0
	temp := t
	for i := 31; i >= 0; i-- {
		value := (num >> i) & 1
		if temp.next[1-value] != nil { // 能取到1
			res = res | (1 << i) // 结果在该位可以为1，该为置为1
			temp = temp.next[1-value]
		} else {
			temp = temp.next[value]
		}
	}
	return res
}
