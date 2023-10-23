package main

import "fmt"

func main() {
	fmt.Println(createSortedArray([]int{1, 5, 6, 2}))
}

// leetcode1649_通过指令创建有序数组
var mod = 1000000007

func createSortedArray(instructions []int) int {
	res := 0
	c = make([]int, 100002)
	length = 100001
	for i := 0; i < len(instructions); i++ {
		value := instructions[i]
		upData(value, 1) // 次数+1
		left := getSum(value - 1)
		right := getSum(100000) - getSum(value)
		res = (res + min(left, right)) % mod
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
