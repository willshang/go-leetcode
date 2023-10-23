package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
}

func reversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	temp := make([]int, 0)
	for i := 0; i < n; i++ {
		// 会存在负数的情况：[-5,-5] => -5 > -10
		// 所以加入 2*nums[i] 参与离散化，后面直接查询2*nums[i]
		temp = append(temp, nums[i], 2*nums[i])
	}
	tempArr, m := getLSH(temp)
	total := len(tempArr)
	arr = make([]int, 4*total+1)
	res := 0
	for i := n - 1; i >= 0; i-- {
		index := m[nums[i]]
		// nums[i] > 2*nums[j] => 求小于nums[i]，其中2*nums[j]已经插入
		count := query(0, 1, total, 1, index-1) // 统计插入了多少个小于nums[i]的数
		res = res + count                       //
		update(0, 1, total, m[2*nums[i]], 1)    // 2*nums[j]次数+1
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
