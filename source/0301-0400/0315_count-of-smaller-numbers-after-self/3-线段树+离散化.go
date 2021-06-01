package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}

func countSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	tempArr, m := getLSH(nums)
	total := len(tempArr)
	arr = make([]int, 4*total+1)
	for i := n - 1; i >= 0; i-- {
		index := m[nums[i]]
		res[i] = query(0, 1, total, 1, index-1)
		update(0, 1, total, index, 1)
	}
	return res
}

func getLSH(a []int) ([]int, map[int]int) {
	n := len(a)
	arr := make([]int, n)
	copy(arr, a)
	sort.Ints(arr) // 排序
	m := make(map[int]int)
	m[arr[0]] = 1
	index := 1
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			index++
			m[arr[i]] = index
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = m[a[i]]
	}
	return res, m
}

var arr []int // 线段树

func update(id int, left, right, x int, value int) {
	if left > x || right < x {
		return
	}
	arr[id] = arr[id] + value
	if left == right {
		return
	}
	mid := left + (right-left)/2
	update(2*id+1, left, mid, x, value)    // 左节点
	update(2*id+2, mid+1, right, x, value) // 右节点
}

func query(id int, left, right, queryLeft, queryRight int) int {
	if left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		return arr[id]
	}
	mid := left + (right-left)/2
	return query(id*2+1, left, mid, queryLeft, queryRight) +
		query(id*2+2, mid+1, right, queryLeft, queryRight)
}
