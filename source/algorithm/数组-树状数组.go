package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	n := len(arr)
	c = make([]int, n+1)
	length = n
	for i := 0; i < n; i++ {
		upData(arr[i], 1)
	}
	for i := 0; i < n; i++ {
		fmt.Println(i, getSum(i))
	}
}

var length int
var c []int // 树状数组

// 取2^k：2^k = i&(-i)
// x&(-x)，
// x为0时结果为0；
// x为奇数时，结果为1；
// x为偶数时，结果为x中2的最大次方的因子。
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
