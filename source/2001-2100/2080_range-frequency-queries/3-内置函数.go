package main

import "sort"

type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = append(m[arr[i]], i)
	}
	return RangeFreqQuery{m: m}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	arr := this.m[value]
	l := sort.SearchInts(arr, left)
	r := sort.SearchInts(arr, right+1)
	return r - l
}
