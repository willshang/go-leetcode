package main

import "fmt"

func main() {
	for i := 1; i < 32; i++ {
		fmt.Println(i, minOperations(i))
	}
}

// leetcode1551_使数组中所有元素相等的最小操作数
func minOperations(n int) int {
	res := 0
	if n == 1 {
		return res
	}
	if n%2 == 1 {
		value := n / 2
		res = value * (value + 1)
	} else {
		value := n / 2
		res = value * value
	}
	return res
}
