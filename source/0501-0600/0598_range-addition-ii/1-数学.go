package main

import "fmt"

func main() {
	m := 3
	n := 3
	operations := [][]int{[]int{2, 2}, []int{3, 3}}
	fmt.Println(maxCount(m, n, operations))
}

/*
因为每一次操作都是从M[0][0]开始的一个矩形，
所以每一次更新要考虑前后两次操作的公共区域，
这个公共区域才是被加一的区域，
最后一定是重复最多的公共区域有最大的数字。
*/
// leetcode598_范围求和II
func maxCount(m int, n int, ops [][]int) int {
	for _, o := range ops {
		m = min(m, o[0])
		n = min(n, o[1])
	}
	return m * n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
