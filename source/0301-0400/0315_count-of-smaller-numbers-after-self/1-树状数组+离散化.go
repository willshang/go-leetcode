package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}

// leetcode315_计算右侧小于当前元素的个数
func countSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	arr, _ := getLSH(nums)
	length = n
	c = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		res[i] = getSum(arr[i] - 1)
		upData(arr[i], 1)
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

var length int
var c []int // 树状数组

func lowBit(x int) int {
	return x & (-x)
}

// 单点修改
func upData(i, k int) { // 在i位置加上k
	for i <= length {
		c[i] = c[i] + k
		i = i + lowBit(i) // i = i + 2^k
	}
}

// 区间查询
func getSum(i int) int {
	res := 0
	for i > 0 {
		res = res + c[i]
		i = i - lowBit(i)
	}
	return res
}
