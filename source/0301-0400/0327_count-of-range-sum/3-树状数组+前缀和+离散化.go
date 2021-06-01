package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
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
	length = len(tempArr)
	c = make([]int, length+1)
	upData(m[0], 1) // 注意0
	res := 0
	for i := 1; i < len(preSum); i++ {
		left := m[preSum[i]-upper]
		right := m[preSum[i]-lower]
		index := m[preSum[i]]
		count := getSum(right) - getSum(left-1)
		res = res + count
		upData(index, 1)
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
