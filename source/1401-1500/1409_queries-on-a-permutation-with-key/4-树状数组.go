package main

import "fmt"

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}

func processQueries(queries []int, m int) []int {
	n := len(queries)
	res := make([]int, n)
	// 使用树状数组
	c = make([]int, n+m+1)
	length = n + m
	arr := make([]int, m+1)
	for i := 1; i <= m; i++ {
		arr[i] = n + i
		upData(n+i, 1) // 数组长度n+m+1，从后面n开始存
	}

	for i := 0; i < n; i++ {
		cur := arr[queries[i]]
		upData(cur, -1)      // cur位置-1
		res[i] = getSum(cur) // cur之前的个数
		cur = n - i
		arr[queries[i]] = cur // 移到到n-i，每次往前移动1位
		upData(cur, 1)        // cur位置+1
	}
	return res
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
