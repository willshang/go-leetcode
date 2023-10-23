package main

import "sort"

// leetcode2080_区间内查询数字的频率
type RangeFreqQuery struct {
	m map[int]sort.IntSlice
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int]sort.IntSlice)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = append(m[arr[i]], i)
	}
	return RangeFreqQuery{m: m}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	arr := this.m[value]
	arr = arr[arr.Search(left):]
	return arr.Search(right + 1)
}
