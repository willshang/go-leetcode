package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxRotateFunction([]int{math.MinInt32, math.MinInt32}))
}

// leetcode396_旋转函数
func maxRotateFunction(A []int) int {
	res := 0
	total := 0
	n := len(A)
	for i := 0; i < n; i++ {
		total = total + A[i]
		res = res + i*A[i]
	}
	temp := res
	for i := 0; i < n; i++ {
		// 最后移动到第一位，其他右移一位
		temp = temp + total      // 每一位都加一次
		temp = temp - n*A[n-1-i] // 最后一位删除
		if temp > res {
			res = temp
		}
	}
	return res
}
