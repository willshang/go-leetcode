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
	arr, m := getLSH(temp)
	length = len(arr)
	c = make([]int, length+1)
	res := 0
	for i := n - 1; i >= 0; i-- {
		index := m[nums[i]]
		// nums[i] > 2*nums[j] => 求小于nums[i]，其中2*nums[j]已经插入
		count := getSum(index - 1) // 统计插入了多少个小于nums[i]的数
		res = res + count          //
		upData(m[2*nums[i]], 1)    // 2*nums[j]次数+1
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
