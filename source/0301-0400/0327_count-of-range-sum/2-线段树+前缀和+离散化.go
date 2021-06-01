package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))

	arr := []int{5, -23, -5, -1, -21, 13, 15, 7, 18, 4, 7, 26, 29, -7, -28, 11, -20, -29, 19, 22, 15, 25, 17, -13, 11, -15,
		19, -8, 3, 12, -1, 2, -1, -21, -10, -7, 14, -12, -14, -8, -1, -30, 19, -27, 16, 2, -15, 23, 6, 14, 23,
		2, -4, 4, -9, -8, 10, 20, -29, 29}
	fmt.Println(countRangeSum(arr, -19, 10))
}

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	preSum := make([]int, n+1) // 前缀和
	temp := make([]int, 0)
	temp = append(temp, 0) // 注意0
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
		temp = append(temp, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}
	tempArr, m := getLSH(temp) // 离散化
	total := len(tempArr)
	arr = make([]int, 4*total+1)
	update(0, 1, total, m[0], 1) // 注意0
	res := 0
	for i := 1; i < len(preSum); i++ {
		left := m[preSum[i]-upper]
		right := m[preSum[i]-lower]
		index := m[preSum[i]]
		count := query(0, 1, total, left, right)
		res = res + count
		update(0, 1, total, index, 1)
	}
	return res
}

// 离散化
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
